/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SchedulerConfiguration {
    #[serde(rename = "CreateIndex", skip_serializing_if = "Option::is_none")]
    pub create_index: Option<i32>,
    #[serde(rename = "MemoryOversubscriptionEnabled", skip_serializing_if = "Option::is_none")]
    pub memory_oversubscription_enabled: Option<bool>,
    #[serde(rename = "ModifyIndex", skip_serializing_if = "Option::is_none")]
    pub modify_index: Option<i32>,
    #[serde(rename = "PreemptionConfig", skip_serializing_if = "Option::is_none")]
    pub preemption_config: Option<Box<crate::models::PreemptionConfig>>,
    #[serde(rename = "SchedulerAlgorithm", skip_serializing_if = "Option::is_none")]
    pub scheduler_algorithm: Option<String>,
}

impl SchedulerConfiguration {
    pub fn new() -> SchedulerConfiguration {
        SchedulerConfiguration {
            create_index: None,
            memory_oversubscription_enabled: None,
            modify_index: None,
            preemption_config: None,
            scheduler_algorithm: None,
        }
    }
}

