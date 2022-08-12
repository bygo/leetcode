# Link: https://leetcode.cn/problems/game-play-analysis-iii

SELECT `a1`.`player_id`         `player_id`,
       `a1`.`event_date`        `event_date`,
       SUM(`a2`.`games_played`) `games_played_so_far`
FROM `activity` `a2`
         JOIN
     `activity` `a1`
     ON `a1`.`player_id` = `a2`.`player_id`
         AND `a2`.`event_date` <= `a1`.`event_date`
GROUP BY `a1`.`player_id`, `a1`.`event_date`