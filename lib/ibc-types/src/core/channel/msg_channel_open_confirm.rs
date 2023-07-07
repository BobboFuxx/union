use crate::{core::client::height::Height, CosmosAccountId, MsgIntoProto};

#[derive(Debug, Clone)]
pub struct MsgChannelOpenConfirm {
    pub port_id: String,
    pub channel_id: String,
    pub proof_ack: Vec<u8>,
    pub proof_height: Height,
}

impl MsgIntoProto for MsgChannelOpenConfirm {
    type Proto = protos::ibc::core::channel::v1::MsgChannelOpenConfirm;

    fn into_proto_with_signer(self, signer: &CosmosAccountId) -> Self::Proto {
        Self::Proto {
            port_id: self.port_id,
            channel_id: self.channel_id,
            proof_ack: self.proof_ack,
            proof_height: Some(self.proof_height.into()),
            signer: signer.to_string(),
        }
    }
}

impl From<MsgChannelOpenConfirm> for contracts::ibc_handler::MsgChannelOpenConfirm {
    fn from(msg: MsgChannelOpenConfirm) -> Self {
        Self {
            port_id: msg.port_id,
            channel_id: msg.channel_id,
            proof_ack: msg.proof_ack.into(),
            proof_height: msg.proof_height.into(),
        }
    }
}