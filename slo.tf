resource "datadog_service_level_objective" "SLO 2" {
  name        = "SLO 2"
  type        = "metric"
  description = "Fancy"
  query {
    numerator   = "blah balh"
    denominator = "blah blah"
  }

  thresholds {
    timeframe = "7d"
    target    = 99.9
    warning   = 99.9
  }

  thresholds {
    timeframe = "30d"
    target    = 99.9
    warning   = 99.9
  }

  tags = ["foo:bar", "baz"]
}

