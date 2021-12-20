# Link: https://leetcode-cn.com/problems/evaluate-boolean-expression

SELECT `e`.`left_operand`,
       `operator`,
       `e`.`right_operand`,
       CASE
           WHEN `v1`.`value` > `v2`.`value` AND `operator` = '>' THEN 'true'
           WHEN `v1`.`value` < `v2`.`value` AND `operator` = '<' THEN 'true'
           WHEN `v1`.`value` = `v2`.`value` AND `operator` = '=' THEN 'true'
           ELSE 'false'
           END `value`
FROM `expressions`        `e`
         JOIN `variables` `v1`
         JOIN `variables` `v2`
              ON `e`.`left_operand` = `v1`.`name` AND `e`.`right_operand` = `v2`.`name`