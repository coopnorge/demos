###############################################################################
# pubsub topic
###############################################################################

resource "google_pubsub_schema" "daeca_cloudevents" {
  name       = "daeca-cloudevent"
  type       = "PROTOCOL_BUFFER"
  definition = file("../spec/proto_standalone/cloudevents_bq.proto")
}

resource "google_pubsub_topic" "daeca_cloudevents" {
  name                       = "daeca-cloudevents"
  message_retention_duration = "${7 * 24 * 60 * 60}s"

  schema_settings {
    schema   = google_pubsub_schema.daeca_cloudevents.id
    encoding = "BINARY"
  }
}

resource "google_pubsub_subscription" "daeca_cloudevents_pull" {
  name                       = "daeca-cloudevents-pull"
  topic                      = google_pubsub_topic.daeca_cloudevents.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
  retry_policy {
    minimum_backoff = "1s"
    maximum_backoff = "60s"
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.daeca_cloudevents_pull_dl.id
    max_delivery_attempts = 10
  }
}

resource "google_pubsub_topic" "daeca_cloudevents_pull_dl" {
  name                       = "daeca-cloudevents-pull-dl"
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

resource "google_pubsub_subscription" "daeca_cloudevents_pull_dl" {
  name                       = "daeca-cloudevents-pull-dl"
  topic                      = google_pubsub_topic.daeca_cloudevents_pull_dl.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

###############################################################################
# bigquery: table
###############################################################################

resource "google_bigquery_table" "daeca_cloudevents" {
  dataset_id = google_bigquery_dataset.daeca.dataset_id
  table_id   = "daeca-cloudevents"

  time_partitioning {
    field = "publish_time"
    type  = "DAY"
  }

  deletion_protection = false

  schema = file("../spec/bigquery/pubsub_bigquery_cloudevents.json")
}

###############################################################################
# bigquery subscription
###############################################################################

resource "google_pubsub_subscription" "daeca_cloudevents_bq" {
  name                       = "daeca-cloudevents-bq"
  topic                      = google_pubsub_topic.daeca_cloudevents.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
  bigquery_config {
    table            = "${google_bigquery_table.daeca_cloudevents.project}.${google_bigquery_table.daeca_cloudevents.dataset_id}.${google_bigquery_table.daeca_cloudevents.table_id}"
    write_metadata   = true
    use_topic_schema = true
  }
  retry_policy {
    minimum_backoff = "1s"
    maximum_backoff = "60s"
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.daeca_cloudevents_pull_dl.id
    max_delivery_attempts = 10
  }
}

resource "google_pubsub_topic" "daeca_cloudevents_bq_dl" {
  name                       = "daeca-cloudevents-bq-dl"
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}

resource "google_pubsub_subscription" "daeca_cloudevents_bq_dl" {
  name                       = "daeca-cloudevents-bq-dl"
  topic                      = google_pubsub_topic.daeca_cloudevents_bq_dl.id
  ack_deadline_seconds       = 600
  message_retention_duration = "${7 * 24 * 60 * 60}s"
}
