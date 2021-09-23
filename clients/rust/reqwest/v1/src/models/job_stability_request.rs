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
pub struct JobStabilityRequest {
    #[serde(rename = "JobID", skip_serializing_if = "Option::is_none")]
    pub job_id: Option<String>,
    #[serde(rename = "JobVersion", skip_serializing_if = "Option::is_none")]
    pub job_version: Option<i32>,
    #[serde(rename = "Namespace", skip_serializing_if = "Option::is_none")]
    pub namespace: Option<String>,
    #[serde(rename = "Region", skip_serializing_if = "Option::is_none")]
    pub region: Option<String>,
    #[serde(rename = "SecretID", skip_serializing_if = "Option::is_none")]
    pub secret_id: Option<String>,
    #[serde(rename = "Stable", skip_serializing_if = "Option::is_none")]
    pub stable: Option<bool>,
}

impl JobStabilityRequest {
    pub fn new() -> JobStabilityRequest {
        JobStabilityRequest {
            job_id: None,
            job_version: None,
            namespace: None,
            region: None,
            secret_id: None,
            stable: None,
        }
    }
}


