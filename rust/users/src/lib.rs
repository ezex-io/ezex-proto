//! This crate exposes the Protobuf messages and gRPC stubs for the “users” service.
//! Buf has already generated `users.rs` and `users.tonic.rs` into this folder.

pub mod users {
    tonic::include_proto!("users");
}
