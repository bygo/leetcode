# Link: https://leetcode-cn.com/problems/number-of-calls-between-two-persons

SELECT `person1`, `person2`, COUNT(1) `call_count`, SUM(`duration`) `total_duration`
FROM (
         SELECT IF(`from_id` > `to_id`, `to_id`, `from_id`) `person1`,
                IF(`from_id` > `to_id`, `from_id`, `to_id`) `person2`,
                `duration`
         FROM `calls`) `c`
GROUP BY `person1`, `person2`;

SELECT `from_id`        `person1`,
       `to_id`          `person2`,
       COUNT(`from_id`) `call_count`,
       SUM(`duration`)  `total_duration`
FROM `calls`
GROUP BY LEAST(`from_id`, `to_id`), GREATEST(`from_id`, `to_id`)