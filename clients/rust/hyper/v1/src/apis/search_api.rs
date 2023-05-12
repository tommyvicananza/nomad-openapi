/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 * Generated by: https://openapi-generator.tech
 */

use std::rc::Rc;
use std::borrow::Borrow;
#[allow(unused_imports)]
use std::option::Option;

use hyper;
use serde_json;
use futures::Future;

use super::{Error, configuration};
use super::request as __internal_request;

pub struct SearchApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> SearchApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> SearchApiClient<C> {
        SearchApiClient {
            configuration,
        }
    }
}

pub trait SearchApi {
    fn get_fuzzy_search(&self, fuzzy_search_request: crate::models::FuzzySearchRequest, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::FuzzySearchResponse, Error = Error<serde_json::Value>>>;
    fn get_search(&self, search_request: crate::models::SearchRequest, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SearchResponse, Error = Error<serde_json::Value>>>;
}

impl<C: hyper::client::Connect>SearchApi for SearchApiClient<C> {
    fn get_fuzzy_search(&self, fuzzy_search_request: crate::models::FuzzySearchRequest, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::FuzzySearchResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/search/fuzzy".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = wait {
            req = req.with_query_param("wait".to_string(), s.to_string());
        }
        if let Some(ref s) = stale {
            req = req.with_query_param("stale".to_string(), s.to_string());
        }
        if let Some(ref s) = prefix {
            req = req.with_query_param("prefix".to_string(), s.to_string());
        }
        if let Some(ref s) = per_page {
            req = req.with_query_param("per_page".to_string(), s.to_string());
        }
        if let Some(ref s) = next_token {
            req = req.with_query_param("next_token".to_string(), s.to_string());
        }
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(fuzzy_search_request);

        req.execute(self.configuration.borrow())
    }

    fn get_search(&self, search_request: crate::models::SearchRequest, region: Option<&str>, namespace: Option<&str>, index: Option<i32>, wait: Option<&str>, stale: Option<&str>, prefix: Option<&str>, x_nomad_token: Option<&str>, per_page: Option<i32>, next_token: Option<&str>) -> Box<dyn Future<Item = crate::models::SearchResponse, Error = Error<serde_json::Value>>> {
        let mut req = __internal_request::Request::new(hyper::Method::Post, "/search".to_string())
            .with_auth(__internal_request::Auth::ApiKey(__internal_request::ApiKey{
                in_header: true,
                in_query: false,
                param_name: "X-Nomad-Token".to_owned(),
            }))
        ;
        if let Some(ref s) = region {
            req = req.with_query_param("region".to_string(), s.to_string());
        }
        if let Some(ref s) = namespace {
            req = req.with_query_param("namespace".to_string(), s.to_string());
        }
        if let Some(ref s) = wait {
            req = req.with_query_param("wait".to_string(), s.to_string());
        }
        if let Some(ref s) = stale {
            req = req.with_query_param("stale".to_string(), s.to_string());
        }
        if let Some(ref s) = prefix {
            req = req.with_query_param("prefix".to_string(), s.to_string());
        }
        if let Some(ref s) = per_page {
            req = req.with_query_param("per_page".to_string(), s.to_string());
        }
        if let Some(ref s) = next_token {
            req = req.with_query_param("next_token".to_string(), s.to_string());
        }
        if let Some(param_value) = index {
            req = req.with_header_param("index".to_string(), param_value.to_string());
        }
        if let Some(param_value) = x_nomad_token {
            req = req.with_header_param("X-Nomad-Token".to_string(), param_value.to_string());
        }
        req = req.with_body_param(search_request);

        req.execute(self.configuration.borrow())
    }

}
