# Link: https://leetcode-cn.com/problems/weather-type-in-each-country

SELECT `country_name`,
       IF(`weather_state_avg` <= 15, 'Cold',
          IF(`weather_state_avg` >= 25, 'Hot', 'Warm')) `weather_type`
FROM (
         SELECT `country_name`,
                AVG(`weather_state`) `weather_state_avg`
         FROM `countries`        `c`
                  JOIN `weather` `w`
                       ON `c`.`country_id` = `w`.`country_id`
         WHERE LEFT(`day`, 7) = '2019-11'
         GROUP BY `country_name`) `t`