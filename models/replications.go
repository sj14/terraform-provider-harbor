package models

var PathReplication = "/replication/policies"

type ReplicationBody struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SrcRegistry struct {
		ID int `json:"id,omitempty"`
	} `json:"src_registry,omitempty"`
	DestRegistry struct {
		ID int `json:"id,omitempty"`
	} `json:"dest_registry,omitempty"`
	DestNamespace string `json:"dest_namespace,omitempty"`
	Trigger       struct {
		Type            string `json:"type,omitempty"`
		TriggerSettings struct {
			Cron string `json:"cron,omitempty"`
		} `json:"trigger_settings,omitempty"`
	} `json:"trigger,omitempty"`
	Enabled  bool                 `json:"enabled,omitempty"`
	Deletion bool                 `json:"deletion,omitempty"`
	Override bool                 `json:"override,omitempty"`
	Filters  []ReplicationFilters `json:"filters,omitempty"`
}

type ReplicationFilters struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
