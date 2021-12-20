# Link: https://leetcode-cn.com/problems/rectangles-area


SELECT `a`.`id`                                                               `p1`,
       `b`.`id`                                                               `p2`,
       ABS((`a`.`x_value` - `b`.`x_value`) * (`a`.`y_value` - `b`.`y_value`)) `area`
FROM `points`          `a`
         JOIN `points` `b`
WHERE `a`.`id` < `b`.`id`
HAVING `area` > 0
ORDER BY `area` DESC, `p1`, `p2`;
