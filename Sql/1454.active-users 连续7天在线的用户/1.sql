# Link: https://leetcode-cn.com/problems/active-users


SELECT `t3`.`id`, `name`
FROM (
         SELECT DISTINCT `id`
         FROM (
                  SELECT `id`, SUBDATE(`login_date`, `r`) `diff`
                  FROM (
                           SELECT `id`,
                                  `login_date`,
                                  ROW_NUMBER() OVER (PARTITION BY `id` ORDER BY `login_date`) `r`
                           FROM (
                                    SELECT DISTINCT *
                                    FROM `logins`) `l`) `t1`
                  GROUP BY `id`, `diff`
                  HAVING COUNT(*) >= 5) `t2`) `t3`
         JOIN `accounts` USING (`id`)
ORDER BY `id`
