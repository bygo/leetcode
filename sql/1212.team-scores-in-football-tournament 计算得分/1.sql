# Link: https://leetcode.cn/problems/team-scores-in-football-tournament


SELECT `t`.`team_id`, `team_name`, IFNULL(SUM(`points`), 0) `num_points`
FROM `teams`                        `t`
         LEFT JOIN (SELECT `host_team`                                `team_id`,
                           IF(`host_goals` > `guest_goals`, 3,
                              IF(`host_goals` = `guest_goals`, 1, 0)) `points`
                    FROM `matches`
                    UNION ALL
                    SELECT `guest_team`                               `team_id`,
                           IF(`host_goals` < `guest_goals`, 3,
                              IF(`host_goals` = `guest_goals`, 1, 0)) `points`
                    FROM `matches`) `m`
                   ON `t`.`team_id` = `m`.`team_id`
GROUP BY `t`.`team_id`
ORDER BY `num_points` DESC, `t`.`team_id`;
