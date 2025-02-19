// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StreamAvroFileFormat struct {
}

type StreamBackfillAll struct {
	/* MySQL data source objects to avoid backfilling. */
	// +optional
	MysqlExcludedObjects *StreamMysqlExcludedObjects `json:"mysqlExcludedObjects,omitempty"`

	/* PostgreSQL data source objects to avoid backfilling. */
	// +optional
	OracleExcludedObjects *StreamOracleExcludedObjects `json:"oracleExcludedObjects,omitempty"`

	/* PostgreSQL data source objects to avoid backfilling. */
	// +optional
	PostgresqlExcludedObjects *StreamPostgresqlExcludedObjects `json:"postgresqlExcludedObjects,omitempty"`
}

type StreamBackfillNone struct {
}

type StreamBigqueryDestinationConfig struct {
	/* The guaranteed data freshness (in seconds) when querying tables created by the stream.
	Editing this field will only affect new tables created in the future, but existing tables
	will not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s. */
	// +optional
	DataFreshness *string `json:"dataFreshness,omitempty"`

	/* A single target dataset to which all data will be streamed. */
	// +optional
	SingleTargetDataset *StreamSingleTargetDataset `json:"singleTargetDataset,omitempty"`

	/* Destination datasets are created so that hierarchy of the destination data objects matches the source hierarchy. */
	// +optional
	SourceHierarchyDatasets *StreamSourceHierarchyDatasets `json:"sourceHierarchyDatasets,omitempty"`
}

type StreamDatasetTemplate struct {
	/* If supplied, every created dataset will have its name prefixed by the provided value.
	The prefix and name will be separated by an underscore. i.e. _. */
	// +optional
	DatasetIdPrefix *string `json:"datasetIdPrefix,omitempty"`

	/* Immutable. Describes the Cloud KMS encryption key that will be used to protect destination BigQuery
	table. The BigQuery Service Account associated with your project requires access to this
	encryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
	See https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information. */
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty"`

	/* The geographic location where the dataset should reside.
	See https://cloud.google.com/bigquery/docs/locations for supported locations. */
	Location string `json:"location"`
}

type StreamDestinationConfig struct {
	/* A configuration for how data should be loaded to Cloud Storage. */
	// +optional
	BigqueryDestinationConfig *StreamBigqueryDestinationConfig `json:"bigqueryDestinationConfig,omitempty"`

	/* Immutable. Destination connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}. */
	DestinationConnectionProfile string `json:"destinationConnectionProfile"`

	/* A configuration for how data should be loaded to Cloud Storage. */
	// +optional
	GcsDestinationConfig *StreamGcsDestinationConfig `json:"gcsDestinationConfig,omitempty"`
}

type StreamDropLargeObjects struct {
}

type StreamExcludeObjects struct {
	/* PostgreSQL schemas on the server. */
	PostgresqlSchemas []StreamPostgresqlSchemas `json:"postgresqlSchemas"`
}

type StreamGcsDestinationConfig struct {
	/* AVRO file format configuration. */
	// +optional
	AvroFileFormat *StreamAvroFileFormat `json:"avroFileFormat,omitempty"`

	/* The maximum duration for which new events are added before a file is closed and a new file is created.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s. */
	// +optional
	FileRotationInterval *string `json:"fileRotationInterval,omitempty"`

	/* The maximum file size to be saved in the bucket. */
	// +optional
	FileRotationMb *int `json:"fileRotationMb,omitempty"`

	/* JSON file format configuration. */
	// +optional
	JsonFileFormat *StreamJsonFileFormat `json:"jsonFileFormat,omitempty"`

	/* Path inside the Cloud Storage bucket to write data to. */
	// +optional
	Path *string `json:"path,omitempty"`
}

type StreamIncludeObjects struct {
	/* PostgreSQL schemas on the server. */
	PostgresqlSchemas []StreamPostgresqlSchemas `json:"postgresqlSchemas"`
}

type StreamJsonFileFormat struct {
	/* Compression of the loaded JSON file. Possible values: ["NO_COMPRESSION", "GZIP"]. */
	// +optional
	Compression *string `json:"compression,omitempty"`

	/* The schema file format along JSON data files. Possible values: ["NO_SCHEMA_FILE", "AVRO_SCHEMA_FILE"]. */
	// +optional
	SchemaFileFormat *string `json:"schemaFileFormat,omitempty"`
}

type StreamMysqlColumns struct {
	/* Column collation. */
	// +optional
	Collation *string `json:"collation,omitempty"`

	/* Column name. */
	// +optional
	Column *string `json:"column,omitempty"`

	/* The MySQL data type. Full data types list can be found here:
	https://dev.mysql.com/doc/refman/8.0/en/data-types.html. */
	// +optional
	DataType *string `json:"dataType,omitempty"`

	/* Column length. */
	// +optional
	Length *int `json:"length,omitempty"`

	/* Whether or not the column can accept a null value. */
	// +optional
	Nullable *bool `json:"nullable,omitempty"`

	/* The ordinal position of the column in the table. */
	// +optional
	OrdinalPosition *int `json:"ordinalPosition,omitempty"`

	/* Whether or not the column represents a primary key. */
	// +optional
	PrimaryKey *bool `json:"primaryKey,omitempty"`
}

type StreamMysqlDatabases struct {
	/* Database name. */
	Database string `json:"database"`

	/* Tables in the database. */
	// +optional
	MysqlTables []StreamMysqlTables `json:"mysqlTables,omitempty"`
}

type StreamMysqlExcludedObjects struct {
	/* MySQL databases on the server. */
	MysqlDatabases []StreamMysqlDatabases `json:"mysqlDatabases"`
}

type StreamMysqlSourceConfig struct {
	/* MySQL objects to exclude from the stream. */
	// +optional
	ExcludeObjects *StreamExcludeObjects `json:"excludeObjects,omitempty"`

	/* MySQL objects to retrieve from the source. */
	// +optional
	IncludeObjects *StreamIncludeObjects `json:"includeObjects,omitempty"`

	/* Maximum number of concurrent backfill tasks. The number should be non negative.
	If not set (or set to 0), the system's default value will be used. */
	// +optional
	MaxConcurrentBackfillTasks *int `json:"maxConcurrentBackfillTasks,omitempty"`

	/* Maximum number of concurrent CDC tasks. The number should be non negative.
	If not set (or set to 0), the system's default value will be used. */
	// +optional
	MaxConcurrentCdcTasks *int `json:"maxConcurrentCdcTasks,omitempty"`
}

type StreamMysqlTables struct {
	/* MySQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything. */
	// +optional
	MysqlColumns []StreamMysqlColumns `json:"mysqlColumns,omitempty"`

	/* Table name. */
	Table string `json:"table"`
}

type StreamOracleColumns struct {
	/* Column name. */
	// +optional
	Column *string `json:"column,omitempty"`

	/* The Oracle data type. Full data types list can be found here:
	https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html. */
	// +optional
	DataType *string `json:"dataType,omitempty"`

	/* Column encoding. */
	// +optional
	Encoding *string `json:"encoding,omitempty"`

	/* Column length. */
	// +optional
	Length *int `json:"length,omitempty"`

	/* Whether or not the column can accept a null value. */
	// +optional
	Nullable *bool `json:"nullable,omitempty"`

	/* The ordinal position of the column in the table. */
	// +optional
	OrdinalPosition *int `json:"ordinalPosition,omitempty"`

	/* Column precision. */
	// +optional
	Precision *int `json:"precision,omitempty"`

	/* Whether or not the column represents a primary key. */
	// +optional
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	/* Column scale. */
	// +optional
	Scale *int `json:"scale,omitempty"`
}

type StreamOracleExcludedObjects struct {
	/* Oracle schemas/databases in the database server. */
	OracleSchemas []StreamOracleSchemas `json:"oracleSchemas"`
}

type StreamOracleSchemas struct {
	/* Tables in the database. */
	// +optional
	OracleTables []StreamOracleTables `json:"oracleTables,omitempty"`

	/* Schema name. */
	Schema string `json:"schema"`
}

type StreamOracleSourceConfig struct {
	/* Configuration to drop large object values. */
	// +optional
	DropLargeObjects *StreamDropLargeObjects `json:"dropLargeObjects,omitempty"`

	/* Oracle objects to exclude from the stream. */
	// +optional
	ExcludeObjects *StreamExcludeObjects `json:"excludeObjects,omitempty"`

	/* Oracle objects to retrieve from the source. */
	// +optional
	IncludeObjects *StreamIncludeObjects `json:"includeObjects,omitempty"`

	/* Maximum number of concurrent backfill tasks. The number should be non negative.
	If not set (or set to 0), the system's default value will be used. */
	// +optional
	MaxConcurrentBackfillTasks *int `json:"maxConcurrentBackfillTasks,omitempty"`

	/* Maximum number of concurrent CDC tasks. The number should be non negative.
	If not set (or set to 0), the system's default value will be used. */
	// +optional
	MaxConcurrentCdcTasks *int `json:"maxConcurrentCdcTasks,omitempty"`

	/* Configuration to drop large object values. */
	// +optional
	StreamLargeObjects *StreamStreamLargeObjects `json:"streamLargeObjects,omitempty"`
}

type StreamOracleTables struct {
	/* Oracle columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything. */
	// +optional
	OracleColumns []StreamOracleColumns `json:"oracleColumns,omitempty"`

	/* Table name. */
	Table string `json:"table"`
}

type StreamPostgresqlColumns struct {
	/* Column name. */
	// +optional
	Column *string `json:"column,omitempty"`

	/* The PostgreSQL data type. Full data types list can be found here:
	https://www.postgresql.org/docs/current/datatype.html. */
	// +optional
	DataType *string `json:"dataType,omitempty"`

	/* Column length. */
	// +optional
	Length *int `json:"length,omitempty"`

	/* Whether or not the column can accept a null value. */
	// +optional
	Nullable *bool `json:"nullable,omitempty"`

	/* The ordinal position of the column in the table. */
	// +optional
	OrdinalPosition *int `json:"ordinalPosition,omitempty"`

	/* Column precision. */
	// +optional
	Precision *int `json:"precision,omitempty"`

	/* Whether or not the column represents a primary key. */
	// +optional
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	/* Column scale. */
	// +optional
	Scale *int `json:"scale,omitempty"`
}

type StreamPostgresqlExcludedObjects struct {
	/* PostgreSQL schemas on the server. */
	PostgresqlSchemas []StreamPostgresqlSchemas `json:"postgresqlSchemas"`
}

type StreamPostgresqlSchemas struct {
	/* Tables in the schema. */
	// +optional
	PostgresqlTables []StreamPostgresqlTables `json:"postgresqlTables,omitempty"`

	/* Database name. */
	Schema string `json:"schema"`
}

type StreamPostgresqlSourceConfig struct {
	/* PostgreSQL objects to exclude from the stream. */
	// +optional
	ExcludeObjects *StreamExcludeObjects `json:"excludeObjects,omitempty"`

	/* PostgreSQL objects to retrieve from the source. */
	// +optional
	IncludeObjects *StreamIncludeObjects `json:"includeObjects,omitempty"`

	/* Maximum number of concurrent backfill tasks. The number should be non
	negative. If not set (or set to 0), the system's default value will be used. */
	// +optional
	MaxConcurrentBackfillTasks *int `json:"maxConcurrentBackfillTasks,omitempty"`

	/* The name of the publication that includes the set of all tables
	that are defined in the stream's include_objects. */
	Publication string `json:"publication"`

	/* The name of the logical replication slot that's configured with
	the pgoutput plugin. */
	ReplicationSlot string `json:"replicationSlot"`
}

type StreamPostgresqlTables struct {
	/* PostgreSQL columns in the schema. When unspecified as part of include/exclude objects, includes/excludes everything. */
	// +optional
	PostgresqlColumns []StreamPostgresqlColumns `json:"postgresqlColumns,omitempty"`

	/* Table name. */
	Table string `json:"table"`
}

type StreamSingleTargetDataset struct {
	/* Dataset ID in the format projects/{project}/datasets/{dataset_id} or
	{project}:{dataset_id}. */
	DatasetId string `json:"datasetId"`
}

type StreamSourceConfig struct {
	/* MySQL data source configuration. */
	// +optional
	MysqlSourceConfig *StreamMysqlSourceConfig `json:"mysqlSourceConfig,omitempty"`

	/* MySQL data source configuration. */
	// +optional
	OracleSourceConfig *StreamOracleSourceConfig `json:"oracleSourceConfig,omitempty"`

	/* PostgreSQL data source configuration. */
	// +optional
	PostgresqlSourceConfig *StreamPostgresqlSourceConfig `json:"postgresqlSourceConfig,omitempty"`

	/* Immutable. Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}. */
	SourceConnectionProfile string `json:"sourceConnectionProfile"`
}

type StreamSourceHierarchyDatasets struct {
	/* Dataset template used for dynamic dataset creation. */
	DatasetTemplate StreamDatasetTemplate `json:"datasetTemplate"`
}

type StreamStreamLargeObjects struct {
}

type DatastreamStreamSpec struct {
	/* Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded. */
	// +optional
	BackfillAll *StreamBackfillAll `json:"backfillAll,omitempty"`

	/* Backfill strategy to disable automatic backfill for the Stream's objects. */
	// +optional
	BackfillNone *StreamBackfillNone `json:"backfillNone,omitempty"`

	/* Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
	will be encrypted using an internal Stream-specific encryption key provisioned through KMS. */
	// +optional
	CustomerManagedEncryptionKey *string `json:"customerManagedEncryptionKey,omitempty"`

	/* Desired state of the Stream. Set this field to 'RUNNING' to start the stream, and 'PAUSED' to pause the stream. */
	// +optional
	DesiredState *string `json:"desiredState,omitempty"`

	/* Destination connection profile configuration. */
	DestinationConfig StreamDestinationConfig `json:"destinationConfig"`

	/* Display name. */
	DisplayName string `json:"displayName"`

	/* Immutable. The name of the location this stream is located in. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The streamId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Source connection profile configuration. */
	SourceConfig StreamSourceConfig `json:"sourceConfig"`
}

type DatastreamStreamStatus struct {
	/* Conditions represent the latest available observations of the
	   DatastreamStream's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The stream's name. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The state of the stream. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatastreamStream is the Schema for the datastream API
// +k8s:openapi-gen=true
type DatastreamStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamStreamSpec   `json:"spec,omitempty"`
	Status DatastreamStreamStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatastreamStreamList contains a list of DatastreamStream
type DatastreamStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamStream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamStream{}, &DatastreamStreamList{})
}
