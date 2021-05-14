# Link: https://leetcode-cn.com/problems/number-of-calls-between-two-persons

SELECT `person1`, `person2`, count(1) `call_count`, sum(`duration`) `total_duration`
FROM (
         SELECT if(`from_id` > `to_id`, `to_id`, `from_id`) `person1`,
                if(`from_id` > `to_id`, `from_id`, `to_id`) `person2`,
                `duration`
         FROM `calls`) `c`
GROUP BY `person1`, `person2`;

SELECT `from_id`        `person1`,
       `to_id`          `person2`,
       count(`from_id`) `call_count`,
       sum(`duration`)  `total_duration`
FROM `calls`
GROUP BY least(`from_id`, `to_id`), greatest(`from_id`, `to_id`)