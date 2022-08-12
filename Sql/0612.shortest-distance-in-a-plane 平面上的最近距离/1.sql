# Link: https://leetcode.cn/problems/shortest-distance-in-a-plane

SELECT ROUND(SQRT(MIN((POW(`p1`.`x` - `p2`.`x`, 2) + POW(`p1`.`y` - `p2`.`y`, 2)))), 2) `shortest`
FROM `point_2d` `p1`
         JOIN
     `point_2d` `p2` ON (`p1`.`x` != `p2`.`x` AND `p1`.`y` = `p2`.`y`)
         OR (`p1`.`x` != `p2`.`x` AND `p1`.`y` != `p2`.`y`)
         OR (`p1`.`x` = `p2`.`x` AND `p1`.`y` != `p2`.`y`)
