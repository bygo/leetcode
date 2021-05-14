# Link: https://leetcode-cn.com/problems/ads-performance

SELECT `ad_id`,
       ROUND(IFNULL(SUM(`action` = 'Clicked') /
                    (SUM(`action` = 'Clicked') + SUM(`action` = 'Viewed')) * 100, 0), 2) `ctr`
FROM `ads`
GROUP BY `ad_id`
ORDER BY `ctr` DESC, `ad_id`