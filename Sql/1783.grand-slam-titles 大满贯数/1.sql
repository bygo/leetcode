# Link: https://leetcode-cn.com/problems/grand-slam-titles


WITH `t` AS (SELECT `wimbledon` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `fr_open` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `us_open` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `au_open` `tournament`
             FROM `championships`)

SELECT `p`.`player_id`, `p`.`player_name`, COUNT(*) `grand_slams_count`
FROM `t`
         LEFT JOIN `players` `p` ON `t`.`tournament` = `p`.`player_id`
GROUP BY `p`.`player_id`, `p`.`player_name`