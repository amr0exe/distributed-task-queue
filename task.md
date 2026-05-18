
# Constraints for this Phase

- create artificial concurrency pressure
  - 1000 concurrent updates to same task

- force connection exhaustion

- force slow queries

- force transaction conflicts

- force restart recovery
  - kill server mid-request

    **Verify**
    .DB survives
    .state consistent

## Priority Order

- Stage1 _ Infra foundation
  - task1: docker-compose with postgresql
  - task2: add goose migration setup
  - task3: initial task table

- Stage2 _ Database Connectivity
  - task1: add pgx connection pool (connection pool, pool config, context usage)
  - task2: health check endpoint

- Stage3 _ Repository Layer
  - task1: Create Repository Interface
  - task2: implement postgres repo for Create and Get only, for now
    - consider
      - context propagation
      - error handling (**differentiate** *[no_rows | connection_failure | query_failure]*)
      - row scanning (*mapping DB rows -> Go structs*)

- Stage4 _ Persistence Validation
  - task1: Verify persistence across restart (*create -> stop -> restart -> retrieve task*)
  - task2: add full crud (*now only introduce update, delete, list*)

- Stage5 _ Introduce pressure
  - task1: add concurrent DB writes
    - hammer api with:
      - go routines
      - curl loops
  - task2: Introduce transactions

- Stage6 _ Performace Awareness
  - task: start looking for indexing, explain analyze, query planning
