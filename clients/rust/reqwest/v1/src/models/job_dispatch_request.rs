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
pub struct JobDispatchRequest {
    #[serde(rename = "IdPrefixTemplate", skip_serializing_if = "Option::is_none")]
    pub id_prefix_template: Option<String>,
    #[serde(rename = "JobID", skip_serializing_if = "Option::is_none")]
    pub job_id: Option<String>,
    #[serde(rename = "Meta", skip_serializing_if = "Option::is_none")]
    pub meta: Option<::std::collections::HashMap<String, String>>,
    #[serde(rename = "Payload", skip_serializing_if = "Option::is_none")]
    pub payload: Option<String>,
}

impl JobDispatchRequest {
    pub fn new() -> JobDispatchRequest {
        JobDispatchRequest {
            id_prefix_template: None,
            job_id: None,
            meta: None,
            payload: None,
        }
    }
}


