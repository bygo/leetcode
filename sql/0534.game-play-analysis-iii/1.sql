# Title: Game Play Analysis III
# Link: https://leetcode-cn.com/problems/game-play-analysis-iii

SELECT `a2`.`player_id`         AS `player_id`,
       `a2`.`event_date`        AS `event_date`,
       SUM(`a1`.`games_played`) AS `games_played_so_far`
FROM `Activity` `a1`,
     `Activity` `a2`
WHERE `a1`.`player_id` = `a2`.`player_id`
  AND `a1`.`event_date` <= `a2`.`event_date`
GROUP BY `a2`.`player_id`, `a2`.`event_date`