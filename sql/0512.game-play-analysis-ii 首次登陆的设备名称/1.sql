# Link: https://leetcode-cn.com/problems/game-play-analysis-ii

SELECT `player_id`, `device_id`
FROM `Activity`
WHERE (`player_id`, `event_date`) IN (SELECT `player_id`, MIN(`event_date`)
                                      FROM `Activity`
                                      GROUP BY `player_id`);

#

SELECT `player_id`, `device_id`
FROM (SELECT `player_id`,
             `device_id`,
             `event_date`,
             MIN(`event_date`) OVER (PARTITION BY `player_id` ) `m`
      FROM `Activity`) `t`
WHERE `m` = `event_date`
