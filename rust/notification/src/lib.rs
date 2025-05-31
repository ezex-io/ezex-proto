pub mod notification {
    // This macro will `include!("notification.rs")` and `include!("notification.tonic.rs")`
    tonic::include_proto!("notification");
}