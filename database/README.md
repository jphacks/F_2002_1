# データベース構成（MySQL）

- pr: primary key, 該当するものには o を書く
- uq: unique key, 該当するものには数字を書く、同じ数字は複合 unique key
- fk: foreign key, テーブルを書く

## users

| name      | Type     | pr  | required | uq  | fk  | default | description |
| --------- | -------- | --- | -------- | --- | --- | ------- | ----------- |
| id        | int      | o   | True     |
| createdAt | Datetime |     | True     |     |
| updatedAt | Datetime |     | True     |     |
| deletedAt | Datetime |     | True     |     |
| email     | string   |     | True     | 1   |
| name      | string   |     | True     |     |

## plants

| name        | Type     | pr  | required | uq  | fk  | default | description |
| ----------- | -------- | --- | -------- | --- | --- | ------- | ----------- |
| id          | int      | o   | True     |
| createdAt   | Datetime |     | True     |     |
| updatedAt   | Datetime |     | True     |     |
| deletedAt   | Datetime |     | True     |     |
| name        | string   |     | True     |     |
| nick_name   | string   |     | True     |     |
| price       | int      |     | True     |     |
| period      | int      |     | True     |     |     |         | 日          |
| difficulty  | int      |     | True     |     |     |         | 1~3?        |
| description | string   |     | True     |     |
| kit_name    | string   |     | True     |     |
| season_from | int      |     | True     |     |     |         | 月          |
| season_to   | int      |     | True     |     |     |         | 月          |

### user_plants

| name               | Type     | pr  | required | uq  | fk     | default            | description      |
| ------------------ | -------- | --- | -------- | --- | ------ | ------------------ | ---------------- |
| id                 | int      | o   | True     |
| createdAt          | Datetime |     | True     |     |
| updatedAt          | Datetime |     | True     |     |
| deletedAt          | Datetime |     | True     |     |
| user_id            | int      |     | True     |     | users  |
| plant_id           | int      |     | True     |     | plants |
| purchase_date      | Date     |     | True     |     |
| start_growing_date | Date     |     | True     |     |
| nick_name          | string   |     | True     |     |        | plant_id.nick_name | ユーザーが変更可 |

### temperatures

| name | Type   | pr  | required | uq  | fk  | default | description  |
| ---- | ------ | --- | -------- | --- | --- | ------- | ------------ |
| id   | int    | o   | True     |
| name | string |     | True     |     |     |         | 暑い、寒い等 |

### plant_temperatures

| name           | Type  | pr  | required | uq  | fk           | default | description |
| -------------- | ----- | --- | -------- | --- | ------------ | ------- | ----------- |
| id             | int   | o   | True     |
| plant_id       | int   |     | True     |     | plants       |
| temperature_id | int   |     | True     |     | temperatures |
| threshold      | float |     | True     |     |              |         | 気温（℃）   |

### waters

| name | Type   | pr  | required | uq  | fk  | default | description    |
| ---- | ------ | --- | -------- | --- | --- | ------- | -------------- |
| id   | int    | o   | True     |
| name | string |     | True     |     |     |         | 多い、少ない等 |

### plant_waters

| name      | Type  | pr  | required | uq  | fk     | default | description      |
| --------- | ----- | --- | -------- | --- | ------ | ------- | ---------------- |
| id        | int   | o   | True     |
| plant_id  | int   |     | True     |     | plants |
| water_id  | int   |     | True     |     | waters |
| threshold | float |     | True     |     |        |         | 水分量を指す数値 |
