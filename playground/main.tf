provider "cloudranger" {
}

resource "cloudranger_backup_policy" "advanced" {
  name = "advanced backup policy"
  description = "advanced backup policy - managed by cloudranger!"
  timezone_location = "Europe/London"
  backup_source = "instances"
  backup_target = "snapshots"
  retention = 24
  create_cron = "50 18 * * *"

  tags = [
    {
      key = "Terraform",
      value = "1"
    }
  ]
}

resource "cloudranger_backup_policy" "database" {
  name = "database backup policy"
  description = "database backup policy for RDS"
  timezone_location = "Europe/London"
  backup_source = "instances"
  backup_target = "snapshots"
  retention = 24
  create_cron = "50 18 * * *"
  tags = [
    {
      key = "Demo",
      value = "1"
    }
  ]
}

resource "cloudranger_schedule" "basic_schedule"{
  name = "basic schedule"
  description = "basic schedule provisioned by terraform"
  timezone_location = "Europe/London"

  schedule = [
    {
      monday = [{ start = 9, end = 18 }],
      tuesday = [{ start = 9, end = 19 }],
      wednesday = [{ start = 9, end = 20 }],
      thursday = [{ start = 9, end = 21 }],
      friday = [
        { start = 9, end = 18 },
        { start = 22, end = 23, action_type = "Reboot" }
      ],
    }
  ]

  active = true
  match_all_tags = false

  tags = [
    {
      key = "workload-type",
      value = "other"
    },
    {
      key = "Terraform",
      value = "1"
    }
  ]
}


resource "cloudranger_schedule" "database"{
  name = "database schedule"
  description = "database schedule provisioned by terraform"
  timezone_location = "Europe/London"

  schedule {
    friday {
      start = 9
      end = 18
    }
  }

  active = true
  match_all_tags = false

  tags {
    key = "Database"
    value = "1"
  }
}
