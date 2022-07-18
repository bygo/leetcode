# Link: https://leetcode.cn/problems/consecutive-numbers

SELECT DISTINCT `num` `consecutivenums`
FROM (SELECT IF(@`pre` = `num`, @`count` := @`count` + 1, @`count` := 1) `counter`,
             @`pre` := `num`                                             `num`
      FROM `logs`,
           (SELECT @`pre` := 0) AS `t`) AS `t`
WHERE 3 <= `counter`;

#

SELECT DISTINCT `num` `consecutivenums`
FROM (
         SELECT *,
                ROW_NUMBER() OVER (PARTITION BY `num` ORDER BY `id`) `rownum`
         FROM `logs`
     ) `t`
GROUP BY (`id` + 1 - `rownum`), `num`
HAVING 3 <= COUNT(*);

#

SELECT DISTINCT `l1`.`num` `consecutivenums`
FROM `logs` `l1`,
     `logs` `l2`,
     `logs` `l3`
WHERE `l1`.`id` = `l2`.`id` - 1
  AND `l2`.`id` = `l3`.`id` - 1
  AND `l1`.`num` = `l2`.`num`
  AND `l2`.`num` = `l3`.`num`