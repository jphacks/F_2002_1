# データベース構成（MySQL）

- pr: primary key, 該当するものには o を書く
- uq: unique key, 該当するものには数字を書く、同じ数字は複合 unique key
- fk: foreign key, テーブルを書く

## plants

| name        | Type     | pr  | required | uq  | fk(onupdate, ondelete) | default | description |
| ----------- | -------- | --- | -------- | --- | ---------------------- | ------- | ----------- |
| id          | int      | o   | True     |
| createdAt   | Datetime |     | True     |
| updatedAt   | Datetime |     | True     |
| deletedAt   | Datetime |     | True     |
| name        | string   |     | True     |
| nick_name   | string   |     | True     |
| price       | int      |     | True     |
| period      | int      |     | True     |     |                        |         | 日          |
| difficulty  | int      |     | True     |     |                        |         | 1~3?        |
| description | string   |     | True     |
| kit_name    | string   |     | True     |
| season_from | int      |     | True     |     |                        |         | 月          |
| season_to   | int      |     | True     |     |                        |         | 月          |

## temperatures

| name | Type   | pr  | required | uq  | fk(onupdate, ondelete) | default | description  |
| ---- | ------ | --- | -------- | --- | ---------------------- | ------- | ------------ |
| id   | int    | o   | True     |
| name | string |     | True     | 1   |                        |         | 暑い、寒い等 |

## plant_temperatures

| name           | Type  | pr  | required | uq  | fk(onupdate, ondelete)         | default | description |
| -------------- | ----- | --- | -------- | --- | ------------------------------ | ------- | ----------- |
| id             | int   | o   | True     |
| plant_id       | int   |     | True     | 1   | plants(cascade, cascade)       |
| temperature_id | int   |     | True     | 1   | temperatures(cascade, cascade) |
| threshold      | float |     | True     |     |                                |         | 気温（℃）   |

## waters

| name | Type   | pr  | required | uq  | fk(onupdate, ondelete) | default | description    |
| ---- | ------ | --- | -------- | --- | ---------------------- | ------- | -------------- |
| id   | int    | o   | True     |
| name | string |     | True     | 1   |                        |         | 多い、少ない等 |

## plant_waters

| name      | Type | pr  | required | uq  | fk(onupdate, ondelete)   | default | description                           |
| --------- | ---- | --- | -------- | --- | ------------------------ | ------- | ------------------------------------- |
| id        | int  | o   | True     |
| plant_id  | int  |     | True     | 1   | plants(cascade, cascade) |
| water_id  | int  |     | True     | 1   | waters(cascade, cascade) |
| threshold | int  |     | True     |     |                          |         | 水分量を指す数値、型は int でよさそう |

## users

| name      | Type     | pr  | required | uq  | fk(onupdate, ondelete) | default | description |
| --------- | -------- | --- | -------- | --- | ---------------------- | ------- | ----------- |
| id        | int      | o   | True     |
| createdAt | Datetime |     | True     |
| updatedAt | Datetime |     | True     |
| deletedAt | Datetime |     | True     |
| uid       | string   |     |          | 1   |
| name      | string   |     | True     |

## cultivations

| name                  | Type     | pr  | required | uq  | fk(onupdate, ondelete)    | default            | description      |
| --------------------- | -------- | --- | -------- | --- | ------------------------- | ------------------ | ---------------- |
| id                    | int      | o   | True     |
| createdAt             | Datetime |     | True     |     |                           |                    | 購入日と同値     |
| updatedAt             | Datetime |     | True     |
| deletedAt             | Datetime |     | True     |
| user_id               | int      |     | True     |     | users(cascade, cascade)   |
| plant_id              | int      |     | True     |     | plants(cascade, restrict) |
| start_cultivating_at  | DateTime |     | False    |
| finish_cultivating_at | DateTime |     | False    |
| nick_name             | string   |     | True     |     |                           | plant_id.nick_name | ユーザーが変更可 |

## waterings

| name           | Type     | pr  | required | uq  | fk(onupdate, ondelete)         | default | description |
| -------------- | -------- | --- | -------- | --- | ------------------------------ | ------- | ----------- |
| id             | int      | o   | True     |
| createdAt      | Datetime |     | True     |
| updatedAt      | Datetime |     | True     |
| deletedAt      | Datetime |     | True     |
| cultivation_id | int      |     | True     |     | cultivations(cascade, cascade) |

## harvestings

| name           | Type     | pr  | required | uq  | fk(onupdate, ondelete)         | default | description |
| -------------- | -------- | --- | -------- | --- | ------------------------------ | ------- | ----------- |
| id             | int      | o   | True     |
| createdAt      | Datetime |     | True     |
| updatedAt      | Datetime |     | True     |
| deletedAt      | Datetime |     | True     |
| cultivation_id | int      |     | True     |     | cultivations(cascade, cascade) |
