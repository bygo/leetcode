# Link: https://leetcode-cn.com/problems/game-play-analysis-v

SELECT `first_day`                 `install_dt`,
       COUNT(DISTINCT `player_id`) `installs`,
       ROUND(
                   (SUM(IF(DATEDIFF(`event_date`, `first_day`) = 1, 1, 0))) / (COUNT(DISTINCT `player_id`)), 2
           )                       `day1_retention`
FROM (
         SELECT `player_id`,
                `event_date`,
                MIN(`event_date`) OVER (PARTITION BY `player_id`) `first_day`
         FROM `activity`
     ) `a`
GROUP BY `first_day`
