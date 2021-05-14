# Link: https://leetcode-cn.com/problems/user-purchase-platform

SELECT `spend_date`,
       `t2`.`platform`,
       sum(if(`t1`.`platform` = `t2`.`platform`, `amount`, 0)) `total_amount`,
       count(if(`t1`.`platform` = `t2`.`platform`, 1, NULL))   `total_users`
FROM (
         SELECT `spend_date`,
                `user_id`,
                if(count(DISTINCT `platform`) = 2, 'both', `platform`) `platform`,
                sum(`amount`)                                          `amount`
         FROM `spending`
         GROUP BY `user_id`, `spend_date`
     ) `t1`
         JOIN
     (
         SELECT 'desktop' `platform`
         UNION
         SELECT 'mobile' `platform`
         UNION
         SELECT 'both' `platform`
     ) `t2`
GROUP BY `spend_date`, `platform`
