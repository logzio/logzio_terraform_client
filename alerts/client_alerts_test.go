package alerts

func createValidAlert() CreateAlertType {
	return CreateAlertType{
		Title:       "this is my title",
		Description: "this is my description",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   OperatorGreaterThan,
		SeverityThresholdTiers: []SeverityThresholdType{
			SeverityThresholdType{
				SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	}
}

func createUpdateAlert() CreateAlertType {
	return CreateAlertType{
		Title:       "this is my updated title",
		Description: "this is my description",
		QueryString: "loglevel:ERROR",
		Filter:      "",
		Operation:   OperatorGreaterThan,
		SeverityThresholdTiers: []SeverityThresholdType{
			SeverityThresholdType{
				SeverityHigh,
				10,
			},
		},
		SearchTimeFrameMinutes:       0,
		NotificationEmails:           []interface{}{},
		IsEnabled:                    true,
		SuppressNotificationsMinutes: 0,
		ValueAggregationType:         AggregationTypeCount,
		ValueAggregationField:        nil,
		GroupByAggregationFields:     []interface{}{"my_field"},
		AlertNotificationEndpoints:   []interface{}{},
	}
}
