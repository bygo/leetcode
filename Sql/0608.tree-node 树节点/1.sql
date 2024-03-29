# Link: https://leetcode.cn/problems/tree-node

SELECT DISTINCT `t1`.`id`,
                IF(`t1`.`p_id` IS NULL, 'Root',
                   IF(`t2`.`id` IS NOT NULL, 'Inner', 'Leaf')) `type`
FROM `tree`               AS `t1`
         LEFT JOIN `tree` AS `t2`
                   ON `t1`.`id` = `t2`.`p_id`