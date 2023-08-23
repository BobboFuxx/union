use serde::{Deserialize, Serialize};
use ssz::{Decode, Encode};

use crate::{
    errors::{MissingField, TryFromBranchError},
    ethereum::H256,
    ethereum_consts_traits::{
        consts::{floorlog2, FINALIZED_ROOT_INDEX, NEXT_SYNC_COMMITTEE_INDEX},
        BYTES_PER_LOGS_BLOOM, MAX_EXTRA_DATA_BYTES, SYNC_COMMITTEE_SIZE,
    },
    ibc::lightclients::ethereum::{
        light_client_header::LightClientHeader, sync_aggregate::SyncAggregate,
        sync_committee::SyncCommittee,
    },
    try_from_proto_branch, Proto, TryFromProtoErrorOf, TypeUrl,
};

/// TODO: Move these to a more central location
pub type NextSyncCommitteeBranch = [H256; floorlog2(NEXT_SYNC_COMMITTEE_INDEX)];
pub type FinalityBranch = [H256; floorlog2(FINALIZED_ROOT_INDEX)];

#[derive(Clone, Debug, PartialEq, Encode, Decode, Serialize, Deserialize)]
#[serde(bound(serialize = "", deserialize = ""))]
pub struct LightClientUpdate<C: SYNC_COMMITTEE_SIZE + BYTES_PER_LOGS_BLOOM + MAX_EXTRA_DATA_BYTES> {
    /// Header attested to by the sync committee
    pub attested_header: LightClientHeader<C>,
    /// Next sync committee corresponding to `attested_header.state_root`
    // TODO: Merge these fields into one
    #[serde(default = "Option::default")]
    pub next_sync_committee: Option<SyncCommittee<C>>,
    #[serde(default)]
    pub next_sync_committee_branch: Option<NextSyncCommitteeBranch>,
    /// Finalized header corresponding to `attested_header.state_root`
    pub finalized_header: LightClientHeader<C>,
    pub finality_branch: FinalityBranch,
    /// Sync committee aggregate signature
    pub sync_aggregate: SyncAggregate<C>,
    /// Slot at which the aggregate signature was created (untrusted)
    #[serde(with = "::serde_utils::string")]
    pub signature_slot: u64,
}

impl<C: SYNC_COMMITTEE_SIZE + BYTES_PER_LOGS_BLOOM + MAX_EXTRA_DATA_BYTES>
    From<LightClientUpdate<C>>
    for protos::union::ibc::lightclients::ethereum::v1::LightClientUpdate
{
    fn from(value: LightClientUpdate<C>) -> Self {
        Self {
            attested_header: Some(value.attested_header.into()),
            next_sync_committee: value.next_sync_committee.map(Into::into),
            next_sync_committee_branch: value
                .next_sync_committee_branch
                .unwrap_or_default()
                .iter()
                .cloned()
                .map(H256::into_bytes)
                .collect(),
            finalized_header: Some(value.finalized_header.into()),
            finality_branch: value
                .finality_branch
                .iter()
                .cloned()
                .map(H256::into_bytes)
                .collect(),
            sync_aggregate: Some(value.sync_aggregate.into()),
            signature_slot: value.signature_slot,
        }
    }
}

#[derive(Debug)]
pub enum TryFromLightClientUpdateError<
    C: SYNC_COMMITTEE_SIZE + BYTES_PER_LOGS_BLOOM + MAX_EXTRA_DATA_BYTES,
> {
    MissingField(MissingField),
    AttestedHeader(TryFromProtoErrorOf<LightClientHeader<C>>),
    NextSyncCommittee(TryFromProtoErrorOf<SyncCommittee<C>>),
    NextSyncCommitteeBranch(TryFromBranchError<NextSyncCommitteeBranch>),
    FinalityBranch(TryFromBranchError<FinalityBranch>),
    SyncAggregate(TryFromProtoErrorOf<SyncAggregate<C>>),
    FinalizedHeader(TryFromProtoErrorOf<LightClientHeader<C>>),
}

impl<C: SYNC_COMMITTEE_SIZE + BYTES_PER_LOGS_BLOOM + MAX_EXTRA_DATA_BYTES>
    TryFrom<protos::union::ibc::lightclients::ethereum::v1::LightClientUpdate>
    for LightClientUpdate<C>
{
    type Error = TryFromLightClientUpdateError<C>;

    fn try_from(
        value: protos::union::ibc::lightclients::ethereum::v1::LightClientUpdate,
    ) -> Result<Self, Self::Error> {
        Ok(Self {
            attested_header: value
                .attested_header
                .ok_or(TryFromLightClientUpdateError::MissingField(MissingField(
                    "attested_header",
                )))?
                .try_into()
                .map_err(TryFromLightClientUpdateError::AttestedHeader)?,
            next_sync_committee: value
                .next_sync_committee
                .map(TryInto::try_into)
                .transpose()
                .map_err(TryFromLightClientUpdateError::NextSyncCommittee)?,
            next_sync_committee_branch: if value.next_sync_committee_branch.is_empty() {
                None
            } else {
                Some(
                    try_from_proto_branch(value.next_sync_committee_branch)
                        .map_err(TryFromLightClientUpdateError::NextSyncCommitteeBranch)?,
                )
            },
            finalized_header: value
                .finalized_header
                .ok_or(TryFromLightClientUpdateError::MissingField(MissingField(
                    "finalized_header",
                )))?
                .try_into()
                .map_err(TryFromLightClientUpdateError::FinalizedHeader)?,
            finality_branch: try_from_proto_branch(value.finality_branch)
                .map_err(TryFromLightClientUpdateError::FinalityBranch)?,
            sync_aggregate: value
                .sync_aggregate
                .ok_or(TryFromLightClientUpdateError::MissingField(MissingField(
                    "sync_aggregate",
                )))?
                .try_into()
                .map_err(TryFromLightClientUpdateError::SyncAggregate)?,
            signature_slot: value.signature_slot,
        })
    }
}

impl TypeUrl for protos::union::ibc::lightclients::ethereum::v1::LightClientUpdate {
    const TYPE_URL: &'static str = "/union.ibc.lightclients.ethereum.v1.LightClientUpdate";
}

impl<C: SYNC_COMMITTEE_SIZE + BYTES_PER_LOGS_BLOOM + MAX_EXTRA_DATA_BYTES> Proto
    for LightClientUpdate<C>
{
    type Proto = protos::union::ibc::lightclients::ethereum::v1::LightClientUpdate;
}