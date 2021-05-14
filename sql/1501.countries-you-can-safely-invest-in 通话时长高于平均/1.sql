# Link: https://leetcode-cn.com/problems/countries-you-can-safely-invest-in


SELECT `country`.`name` `country`
FROM `calls`
         JOIN
     `person`
         JOIN
     `country`
WHERE (`calls`.`caller_id` = `id` OR `calls`.`callee_id` = `id`)
  AND `country`.`country_code` = LEFT(`person`.`phone_number`, 3)
GROUP BY `country`.`country_code`
HAVING AVG(`calls`.`duration`) > (SELECT AVG(`calls`.`duration`) FROM `calls`);
