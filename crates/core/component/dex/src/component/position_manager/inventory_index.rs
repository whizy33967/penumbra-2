use cnidarium::StateWrite;

use crate::{
    lp::position::{self},
    state_key::eviction_queue,
    DirectedTradingPair,
};

pub(crate) trait PositionByInventoryIndex: StateWrite {
    fn index_position_by_inventory(&mut self, position: &position::Position, id: &position::Id) {
        tracing::debug!("indexing position by inventory");
        let canonical_pair = position.phi.pair;
        // A position is bound to an unordered trading pair: A <> B.
        // We want to index the position by inventory for each direction:
        // A -> B
        let pair_ab = DirectedTradingPair::new(canonical_pair.asset_1(), canonical_pair.asset_2());
        let inventory_a = position
            .reserves_for(pair_ab.start)
            .expect("the directed trading pair is correct");
        let key_ab = eviction_queue::inventory_index::key(&pair_ab, inventory_a, id).to_vec();
        self.nonverifiable_put_raw(key_ab, vec![]);

        // B -> A
        let pair_ba = pair_ab.flip();
        let inventory_b = position
            .reserves_for(pair_ba.start)
            .expect("the directed trading pair is correct");
        let key_ba = eviction_queue::inventory_index::key(&pair_ba, inventory_b, id).to_vec();
        self.nonverifiable_put_raw(key_ba, vec![]);
    }

    fn deindex_position_by_inventory(
        &mut self,
        prev_position: &position::Position,
        id: &position::Id,
    ) {
        let canonical_pair = prev_position.phi.pair;

        // To deindex the position, we need to reconstruct the tuple of keys
        // that correspond to each direction of the trading pair:
        // A -> B
        let pair_ab = DirectedTradingPair::new(canonical_pair.asset_1(), canonical_pair.asset_2());
        let inventory_a = prev_position
            .reserves_for(pair_ab.start)
            .expect("the directed trading pair is correct");
        let key_ab = eviction_queue::inventory_index::key(&pair_ab, inventory_a, id).to_vec();
        self.nonverifiable_delete(key_ab);

        // B -> A
        let pair_ba = pair_ab.flip();
        let inventory_b = prev_position
            .reserves_for(pair_ba.start)
            .expect("the directed trading pair is correct");
        let key_ba = eviction_queue::inventory_index::key(&pair_ba, inventory_b, id).to_vec();
        self.nonverifiable_delete(key_ba);
    }
}

impl<T: StateWrite + ?Sized> PositionByInventoryIndex for T {}