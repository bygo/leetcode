# Link: https://leetcode.cn/problems/game-play-analysis-iv

SELECT ROUND(
               COUNT(`a2`.`player_id`) / COUNT(`a1`.`player_id`),
               2) `fraction`
FROM (
         SELECT `player_id`, MIN(`event_date`) `event_date`
         FROM `activity`
         GROUP BY `player_id`) `a1`
         LEFT JOIN `activity`  `a2`
                   ON `a1`.`player_id` = `a2`.`player_id`
                       AND DATEDIFF(`a2`.`event_date`, `a1`.`event_date`) = 1;