# Title: Game Play Analysis I
# Link: https://leetcode-cn.com/problems/game-play-analysis-i

SELECT `player_id`, `device_id`
FROM `Activity`
ORDER BY `event_date` GROUP BY `player_id`
