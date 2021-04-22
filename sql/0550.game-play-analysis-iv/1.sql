# Title: Game Play Analysis IV
# Link: https://leetcode-cn.com/problems/game-play-analysis-iv

SELECT ROUND(COUNT(`a2`.`player_id`) / COUNT(`a1`.`player_id`), 2) AS `fraction`
FROM (SELECT `player_id`, MIN(`event_date`) AS `event_date`
      FROM `Activity`
      GROUP BY `player_id`) `a1`
         LEFT JOIN `Activity` `a2`
                   ON `a1`.`player_id` = `a2`.`player_id`
                       AND DATEDIFF(`a2`.`event_date`, `a1`.`event_date`) = 1;