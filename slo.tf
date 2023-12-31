resource "datadog_service_level_objective" "" {
  name        = ""
  type        = "metric"
  description = ""
  query {
    numerator   = ""
    denominator = ""
  }

  thresholds {
    timeframe = "7d"
    target    = 
    warning   = 
  }

  thresholds {
    timeframe = "30d"
    target    = 
    warning   = 
  }

  tags = ["foo:bar", "baz"]
}

