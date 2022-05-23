## example settings

```yaml
limits_trigger: true		
water_level_limit: 20
water_amount_limit: 0.3
moist_limit: 27.2
scheduled_trigger: false
hour_range: null
location: Frýdek-Místek
irrigation_duration: 10
chart_type: true
language: false
theme: false
lat: 49.6722638
lon: 18.3961799
default_water_amount: 5
```

## example graphql queries

```graphql
mutation {
  createSettings(input: {
    limits_trigger: false,
    water_level_limit: 0,
    water_amount_limit: 0,
    moist_limit: 0,
    scheduled_trigger: false,
    hour_range: 0,
    location: "kokotov",
    irrigation_duration: 0,
    chart_type: false,
    language: false,
    theme: false,
    lat: 0,
    lon: 0,
  }) {
    id
  }
}
mutation {
  updateSettings(input: {
    limits_trigger: true,
    water_level_limit: 0,
    water_amount_limit: 0,
    moist_limit: 0,
    scheduled_trigger: false,
    hour_range: 0,
    location: "kokotov",
    irrigation_duration: 0,
    chart_type: false,
    language: false,
    theme: false,
    lat: 0,
    lon: 0,
  }) {
    id
  }
}

```