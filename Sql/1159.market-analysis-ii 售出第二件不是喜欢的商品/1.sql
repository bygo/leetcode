# Link: https://leetcode-cn.com/problems/market-analysis-ii


SELECT `user_id`                                                                        `seller_id`,
       IF(`t`.`item_brand` != `u`.`favorite_brand` || `seller_id` IS NULL, 'no', 'yes') `2nd_item_fav_brand`
FROM `users`                                                        `u`
         LEFT JOIN (SELECT `seller_id`,
                           `item_brand`,
                           RANK() OVER (PARTITION BY `seller_id` ORDER BY `order_date`) `r`
                    FROM `orders`         AS `o`
                             JOIN `items` AS `i`
                                  ON `o`.`item_id` = `i`.`item_id`) `t`
                   ON `u`.`user_id` = `t`.`seller_id`
                       AND `r` = 2;