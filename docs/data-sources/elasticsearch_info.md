---
subcategory: "Cluster"
layout: ""
page_title: "Elasticstack: elasticstack_elasticsearch_info Data Source"
description: |-
  Gets information about the Elasticsearch cluster.
---

# Data Source: elasticstack_elasticsearch_info

This data source provides the information about the configured Elasticsearch cluster

## Example Usage

```terraform
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_info" "cluster_info" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `cluster_name` (String) Name of the cluster, based on the Cluster name setting setting.
- `cluster_uuid` (String) Unique identifier for the cluster.
- `id` (String) The ID of this resource.
- `name` (String) Name of the node.
- `tagline` (String) Elasticsearh tag line.
- `version` (List of Object) Contains statistics about the number of nodes selected by the request's node filters. (see [below for nested schema](#nestedatt--version))

<a id="nestedatt--version"></a>
### Nested Schema for `version`

Read-Only:

- `build_date` (String)
- `build_flavor` (String)
- `build_hash` (String)
- `build_snapshot` (Boolean)
- `build_type` (String)
- `lucene_version` (String)
- `minimum_index_compatibility_version` (String)
- `minimum_wire_compatibility_version` (String)
- `number` (String)