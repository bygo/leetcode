# Link: https://leetcode-cn.com/problems/tournament-winners

SELECT `group_id`, `player_id`
FROM (
         SELECT `group_id`, `player_id`, SUM(`score`) `score`
         FROM (
                  -- 每个用户总的 first_score
                  SELECT `p1`.`group_id`, p1.`player_id`, SUM(`m1`.`first_score`) `score`
                  FROM `Players` `p1`
                           JOIN `Matches` m1 ON `p1`.`player_id` = `m1`.`first_player`
                  GROUP BY `p1`.`player_id`

                  UNION ALL

                  -- 每个用户总的 second_score
                  SELECT `p2`.`group_id`, `p2`.`player_id`, SUM(`m2`.`second_score`) `score`
                  FROM `Players` `p2`
                           JOIN `Matches` m2 ON `p2`.`player_id` = `m2`.`second_player`
                  GROUP BY `p2`.`player_id`
              ) `s`
         GROUP BY `player_id`
         ORDER BY `score` DESC, `player_id`
     ) `t`
GROUP BY `group_id`
ORDER BY `group_id`