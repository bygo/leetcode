# Link: https://leetcode.cn/problems/shortest-distance-in-a-line

SELECT MIN(ABS(`p1`.`x` - `p2`.`x`)) `shortest`
FROM `point`          `p1`
         JOIN `point` `p2` ON `p1`.`x` != `p2`.`x`;
