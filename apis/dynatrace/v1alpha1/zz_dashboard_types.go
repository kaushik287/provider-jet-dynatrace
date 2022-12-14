/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BoundsObservation struct {
}

type BoundsParameters struct {

	// the height of the tile, in pixels
	// +kubebuilder:validation:Required
	Height *float64 `json:"height" tf:"height,omitempty"`

	// the horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	// +kubebuilder:validation:Required
	Left *float64 `json:"left" tf:"left,omitempty"`

	// the vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	// +kubebuilder:validation:Required
	Top *float64 `json:"top" tf:"top,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// the width of the tile, in pixels
	// +kubebuilder:validation:Required
	Width *float64 `json:"width" tf:"width,omitempty"`
}

type ChartConfigObservation struct {
}

type ChartConfigParameters struct {

	// The optional custom y-axis limits
	// +kubebuilder:validation:Optional
	AxisLimits map[string]*float64 `json:"axisLimits,omitempty" tf:"axis_limits,omitempty"`

	// Either one of `Bit`, `BitPerHour`, `BitPerMinute`, `BitPerSecond`, `Byte`, `BytePerHour`, `BytePerMinute`, `BytePerSecond`, `Cores`, `Count`, `Day`, `DecibelMilliWatt`, `GibiByte`, `Giga`, `GigaByte`, `Hour`, `KibiByte`, `KibiBytePerHour`, `KibiBytePerMinute`, `KibiBytePerSecond`, `Kilo`, `KiloByte`, `KiloBytePerHour`, `KiloBytePerMinute`, `KiloBytePerSecond`, `MebiByte`, `MebiBytePerHour`, `MebiBytePerMinute`, `MebiBytePerSecond`, `Mega`, `MegaByte`, `MegaBytePerHour`, `MegaBytePerMinute`, `MegaBytePerSecond`, `MicroSecond`, `MilliCores`, `MilliSecond`, `MilliSecondPerMinute`, `Minute`, `Month`, `NanoSecond`, `NanoSecondPerMinute`, `NotApplicable`, `PerHour`, `PerMinute`, `PerSecond`, `Percent`, `Pixel`, `Promille`, `Ratio`, `Second`, `State`, `Unspecified`, `Week`, `Year`
	// +kubebuilder:validation:Optional
	LeftAxisCustomUnit *string `json:"leftAxisCustomUnit,omitempty" tf:"left_axis_custom_unit,omitempty"`

	// Defines if a legend should be shown
	// +kubebuilder:validation:Optional
	Legend *bool `json:"legend,omitempty" tf:"legend,omitempty"`

	// Additional information about charted metric
	// +kubebuilder:validation:Optional
	ResultMetadata []ResultMetadataParameters `json:"resultMetadata,omitempty" tf:"result_metadata,omitempty"`

	// Either one of `Bit`, `BitPerHour`, `BitPerMinute`, `BitPerSecond`, `Byte`, `BytePerHour`, `BytePerMinute`, `BytePerSecond`, `Cores`, `Count`, `Day`, `DecibelMilliWatt`, `GibiByte`, `Giga`, `GigaByte`, `Hour`, `KibiByte`, `KibiBytePerHour`, `KibiBytePerMinute`, `KibiBytePerSecond`, `Kilo`, `KiloByte`, `KiloBytePerHour`, `KiloBytePerMinute`, `KiloBytePerSecond`, `MebiByte`, `MebiBytePerHour`, `MebiBytePerMinute`, `MebiBytePerSecond`, `Mega`, `MegaByte`, `MegaBytePerHour`, `MegaBytePerMinute`, `MegaBytePerSecond`, `MicroSecond`, `MilliCores`, `MilliSecond`, `MilliSecondPerMinute`, `Minute`, `Month`, `NanoSecond`, `NanoSecondPerMinute`, `NotApplicable`, `PerHour`, `PerMinute`, `PerSecond`, `Percent`, `Pixel`, `Promille`, `Ratio`, `Second`, `State`, `Unspecified`, `Week`, `Year`
	// +kubebuilder:validation:Optional
	RightAxisCustomUnit *string `json:"rightAxisCustomUnit,omitempty" tf:"right_axis_custom_unit,omitempty"`

	// A list of charted metrics
	// +kubebuilder:validation:Optional
	Series []SeriesParameters `json:"series,omitempty" tf:"series,omitempty"`

	// The type of the chart
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type ConfigObservation struct {
}

type ConfigParameters struct {

	// The color of the metric in the chart, hex format
	// +kubebuilder:validation:Optional
	CustomColor *string `json:"customColor,omitempty" tf:"custom_color,omitempty"`

	// A generated key by the Dynatrace Server
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The timestamp of the last metadata modification, in UTC milliseconds
	// +kubebuilder:validation:Optional
	LastModified *float64 `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type DashboardMetadataObservation struct {
}

type DashboardMetadataParameters struct {

	// Dashboard filter configuration of a dashboard
	// +kubebuilder:validation:Optional
	DynamicFilters []DynamicFiltersParameters `json:"dynamicFilters,omitempty" tf:"dynamic_filters,omitempty"`

	// Global filter Settings for the Dashboard
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// the name of the dashboard
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// the owner of the dashboard
	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// the dashboard is shared (`true`) or private (`false`)
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// represents sharing configuration of a dashboard
	// +kubebuilder:validation:Optional
	SharingDetails []SharingDetailsParameters `json:"sharingDetails,omitempty" tf:"sharing_details,omitempty"`

	// a set of tags assigned to the dashboard
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// a set of all possible global dashboard filters that can be applied to dashboard
	// +kubebuilder:validation:Optional
	ValidFilterKeys []*string `json:"validFilterKeys,omitempty" tf:"valid_filter_keys,omitempty"`
}

type DashboardObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DashboardParameters struct {

	// contains parameters of a dashboard
	// +kubebuilder:validation:Optional
	DashboardMetadata []DashboardMetadataParameters `json:"dashboardMetadata,omitempty" tf:"dashboard_metadata,omitempty"`

	// `metadata` exists for backwards compatibility but shouldn't get specified anymore
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// the tiles this Dashboard consist of
	// +kubebuilder:validation:Optional
	Tile []TileParameters `json:"tile,omitempty" tf:"tile,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type DimensionObservation struct {
}

type DimensionParameters struct {

	// +kubebuilder:validation:Optional
	EntityDimension *bool `json:"entityDimension,omitempty" tf:"entity_dimension,omitempty"`

	// The ID of the dimension by which the metric is split
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The name of the dimension by which the metric is split
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// The splitting value
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type DynamicFiltersObservation struct {
}

type DynamicFiltersParameters struct {

	// A set of all possible global dashboard filters that can be applied to a dashboard
	//
	// Currently supported values are:
	//
	// OS_TYPE,
	// SERVICE_TYPE,
	// DEPLOYMENT_TYPE,
	// APPLICATION_INJECTION_TYPE,
	// PAAS_VENDOR_TYPE,
	// DATABASE_VENDOR,
	// HOST_VIRTUALIZATION_TYPE,
	// HOST_MONITORING_MODE,
	// KUBERNETES_CLUSTER,
	// RELATED_CLOUD_APPLICATION,
	// RELATED_NAMESPACE,
	// TAG_KEY:<tagname>
	// +kubebuilder:validation:Required
	Filters []*string `json:"filters" tf:"filters,omitempty"`

	// A set of entities applied for tag filter suggestions. You can fetch the list of possible values with the [GET all entity types](https://dt-url.net/dw03s7h)request.
	//
	// Only applicable if the **filters** set includes `TAG_KEY:<tagname>`
	// +kubebuilder:validation:Optional
	TagSuggestionTypes []*string `json:"tagSuggestionTypes,omitempty" tf:"tag_suggestion_types,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FilterConfigObservation struct {
}

type FilterConfigParameters struct {

	// Configuration of a custom chart
	// +kubebuilder:validation:Optional
	ChartConfig []ChartConfigParameters `json:"chartConfig,omitempty" tf:"chart_config,omitempty"`

	// The name of the tile, set by user
	// +kubebuilder:validation:Required
	CustomName *string `json:"customName" tf:"custom_name,omitempty"`

	// The default name of the tile
	// +kubebuilder:validation:Required
	DefaultName *string `json:"defaultName" tf:"default_name,omitempty"`

	// Configuration of a custom chart
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// The type of the filter. Possible values are `ALB`, `APPLICATION`, `APPLICATION_METHOD`, `APPMON`, `ASG`, `AWS_CREDENTIALS`, `AWS_CUSTOM_SERVICE`, `AWS_LAMBDA_FUNCTION`, `CLOUD_APPLICATION`, `CLOUD_APPLICATION_INSTANCE`, `CLOUD_APPLICATION_NAMESPACE`, `CONTAINER_GROUP_INSTANCE`, `CUSTOM_APPLICATION`, `CUSTOM_DEVICES`, `CUSTOM_SERVICES`, `DATABASE`, `DATABASE_KEY_REQUEST`, `DCRUM_APPLICATION`, `DCRUM_ENTITY`, `DYNAMO_DB`, `EBS`, `EC2`, `ELB`, `ENVIRONMENT`, `ESXI`, `EXTERNAL_SYNTHETIC_TEST`, `GLOBAL_BACKGROUND_ACTIVITY`, `HOST`, `IOT`, `KUBERNETES_CLUSTER`, `KUBERNETES_NODE`, `MDA_SERVICE`, `MIXED`, `MOBILE_APPLICATION`, `MONITORED_ENTITY`, `NLB`, `PG_BACKGROUND_ACTIVITY`, `PROBLEM`, `PROCESS_GROUP_INSTANCE`, `RDS`, `REMOTE_PLUGIN`, `SERVICE`, `SERVICE_KEY_REQUEST`, `SYNTHETIC_BROWSER_MONITOR`, `SYNTHETIC_HTTPCHECK`, `SYNTHETIC_HTTPCHECK_STEP`, `SYNTHETIC_LOCATION`, `SYNTHETIC_TEST`, `SYNTHETIC_TEST_STEP`, `UI_ENTITY`, `VIRTUAL_MACHINE`, `WEB_CHECK`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FilterManagementZoneObservation struct {
}

type FilterManagementZoneParameters struct {

	// a short description of the Dynatrace entity
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the ID of the Dynatrace entity
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// the name of the Dynatrace entity
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// the management zone this dashboard applies to
	// +kubebuilder:validation:Optional
	ManagementZone []ManagementZoneParameters `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// the default timeframe of the dashboard
	// +kubebuilder:validation:Optional
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type FiltersFilterObservation struct {
}

type FiltersFilterParameters struct {

	// The entity type (e.g. HOST, SERVICE, ...)
	// +kubebuilder:validation:Required
	EntityType *string `json:"entityType" tf:"entity_type,omitempty"`

	// the tiles this Dashboard consist of
	// +kubebuilder:validation:Optional
	Match []MatchParameters `json:"match,omitempty" tf:"match,omitempty"`
}

type FiltersObservation struct {
}

type FiltersParameters struct {

	// the tiles this Dashboard consist of
	// +kubebuilder:validation:Optional
	Filter []FiltersFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`
}

type ManagementZoneObservation struct {
}

type ManagementZoneParameters struct {

	// a short description of the Dynatrace entity
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the ID of the Dynatrace entity
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// the name of the Dynatrace entity
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type MatchObservation struct {
}

type MatchParameters struct {

	// The entity type (e.g. HOST, SERVICE, ...)
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// the tiles this Dashboard consist of
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MetadataObservation struct {
}

type MetadataParameters struct {

	// Dynatrace server version
	// +kubebuilder:validation:Optional
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// A Sorted list of the version numbers of the configuration
	// +kubebuilder:validation:Optional
	ConfigurationVersions []*float64 `json:"configurationVersions,omitempty" tf:"configuration_versions,omitempty"`

	// A Sorted list of the version numbers of the configuration
	// +kubebuilder:validation:Optional
	CurrentConfigurationVersions []*string `json:"currentConfigurationVersions,omitempty" tf:"current_configuration_versions,omitempty"`
}

type ResultMetadataObservation struct {
}

type ResultMetadataParameters struct {

	// Additional metadata for charted metric
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`
}

type SeriesObservation struct {
}

type SeriesParameters struct {

	// The charted aggregation of the metric
	// +kubebuilder:validation:Required
	Aggregation *string `json:"aggregation" tf:"aggregation,omitempty"`

	// +kubebuilder:validation:Optional
	AggregationRate *string `json:"aggregationRate,omitempty" tf:"aggregation_rate,omitempty"`

	// Configuration of the charted metric splitting
	// +kubebuilder:validation:Optional
	Dimension []DimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// The visualization of the timeseries chart
	// +kubebuilder:validation:Required
	EntityType *string `json:"entityType" tf:"entity_type,omitempty"`

	// The name of the charted metric
	// +kubebuilder:validation:Required
	Metric *string `json:"metric" tf:"metric,omitempty"`

	// The charted percentile. Only applicable if the **aggregation** is set to `PERCENTILE`
	// +kubebuilder:validation:Optional
	Percentile *float64 `json:"percentile,omitempty" tf:"percentile,omitempty"`

	// Sort ascending (`true`) or descending (`false`)
	// +kubebuilder:validation:Optional
	SortAscending *bool `json:"sortAscending,omitempty" tf:"sort_ascending,omitempty"`

	// Sort the column (`true`) or (`false`)
	// +kubebuilder:validation:Optional
	SortColumn *bool `json:"sortColumn,omitempty" tf:"sort_column,omitempty"`

	// The visualization of the timeseries chart. Possible values are `AREA`, `BAR` and `LINE`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type SharingDetailsObservation struct {
}

type SharingDetailsParameters struct {

	// If `true`, the dashboard is shared via link and authenticated users with the link can view
	// +kubebuilder:validation:Optional
	LinkShared *bool `json:"linkShared,omitempty" tf:"link_shared,omitempty"`

	// If `true`, the dashboard is published to anyone on this environment
	// +kubebuilder:validation:Optional
	Published *bool `json:"published,omitempty" tf:"published,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type TileFilterObservation struct {
}

type TileFilterParameters struct {

	// the management zone this tile applies to
	// +kubebuilder:validation:Optional
	ManagementZone []FilterManagementZoneParameters `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// the default timeframe of the tile
	// +kubebuilder:validation:Optional
	Timeframe *string `json:"timeframe,omitempty" tf:"timeframe,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

type TileObservation struct {
}

type TileParameters struct {

	// The list of Dynatrace entities, assigned to the tile
	// +kubebuilder:validation:Optional
	AssignedEntities []*string `json:"assignedEntities,omitempty" tf:"assigned_entities,omitempty"`

	// the position and size of a tile
	// +kubebuilder:validation:Optional
	Bounds []BoundsParameters `json:"bounds,omitempty" tf:"bounds,omitempty"`

	// +kubebuilder:validation:Optional
	ChartVisible *bool `json:"chartVisible,omitempty" tf:"chart_visible,omitempty"`

	// The tile is configured and ready to use (`true`) or just placed on the dashboard (`false`)
	// +kubebuilder:validation:Optional
	Configured *bool `json:"configured,omitempty" tf:"configured,omitempty"`

	// The name of the tile, set by user
	// +kubebuilder:validation:Optional
	CustomName *string `json:"customName,omitempty" tf:"custom_name,omitempty"`

	// Include (`false') or exclude (`true`) maintenance windows from availability calculations
	// +kubebuilder:validation:Optional
	ExcludeMaintenanceWindows *bool `json:"excludeMaintenanceWindows,omitempty" tf:"exclude_maintenance_windows,omitempty"`

	// is filter applied to a tile. It overrides dashboard's filter
	// +kubebuilder:validation:Optional
	Filter []TileFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// the position and size of a tile
	// +kubebuilder:validation:Optional
	FilterConfig []FilterConfigParameters `json:"filterConfig,omitempty" tf:"filter_config,omitempty"`

	// The limit of the results, if not set will use the default value of the system
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// The markdown-formatted content of the tile
	// +kubebuilder:validation:Optional
	Markdown *string `json:"markdown,omitempty" tf:"markdown,omitempty"`

	// The metric assigned to the tile
	// +kubebuilder:validation:Optional
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// the name of the tile
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The size of the tile name. Possible values are `small`, `medium` and `large`.
	// +kubebuilder:validation:Optional
	NameSize *string `json:"nameSize,omitempty" tf:"name_size,omitempty"`

	// A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// the type of the tile. Must be either `APPLICATION_WORLDMAP`, `RESOURCES`, `THIRD_PARTY_MOST_ACTIVE`, `UEM_CONVERSIONS_PER_GOAL`, `PROCESS_GROUPS_ONE` or `HOST` .
	// +kubebuilder:validation:Required
	TileType *string `json:"tileType" tf:"tile_type,omitempty"`

	// The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift
	// +kubebuilder:validation:Optional
	TimeFrameShift *string `json:"timeFrameShift,omitempty" tf:"time_frame_shift,omitempty"`

	// The attribute `type` exists for backwards compatibilty. Usage is discouraged. You should use `visualization` instead.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`

	// The visualization of the tile. Possible values are: `COLUMN_CHART`, `FUNNEL`, `LINE_CHART`, `PIE_CHART`, `SINGLE_VALUE`, `TABLE`
	// +kubebuilder:validation:Optional
	Visualization *string `json:"visualization,omitempty" tf:"visualization,omitempty"`

	// Configuration of a User session query visualization tile
	// +kubebuilder:validation:Optional
	VisualizationConfig []VisualizationConfigParameters `json:"visualizationConfig,omitempty" tf:"visualization_config,omitempty"`
}

type VisualizationConfigObservation struct {
}

type VisualizationConfigParameters struct {

	// The axis bucketing when enabled groups similar series in the same virtual axis
	// +kubebuilder:validation:Optional
	HasAxisBucketing *bool `json:"hasAxisBucketing,omitempty" tf:"has_axis_bucketing,omitempty"`

	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +kubebuilder:validation:Optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns,omitempty"`
}

// DashboardSpec defines the desired state of Dashboard
type DashboardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardParameters `json:"forProvider"`
}

// DashboardStatus defines the observed state of Dashboard.
type DashboardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dashboard is the Schema for the Dashboards API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatracejet}
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardSpec   `json:"spec"`
	Status            DashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardList contains a list of Dashboards
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

// Repository type metadata.
var (
	Dashboard_Kind             = "Dashboard"
	Dashboard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dashboard_Kind}.String()
	Dashboard_KindAPIVersion   = Dashboard_Kind + "." + CRDGroupVersion.String()
	Dashboard_GroupVersionKind = CRDGroupVersion.WithKind(Dashboard_Kind)
)

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
