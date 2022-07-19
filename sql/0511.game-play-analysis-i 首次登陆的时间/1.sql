# Link: https://leetcode.cn/problems/game-play-analysis-i

SELECT `player_id`,
       MIN(`event_date`) `first_login`
FROM `activity`
GROUP BY `player_id`