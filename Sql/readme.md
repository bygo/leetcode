# 0175.combine-two-tables 给所有人加上地址 
```sql
# Link: https://leetcode-cn.com/problems/combine-two-tables

SELECT `firstname`, `lastname`, `city`, `state`
FROM `person`
         LEFT JOIN `address` ON `person`.`personid` = `address`.`personid`
```

# 0176.second-highest-salary 找第二名 Null 也返回 
```sql
# Link: https://leetcode-cn.com/problems/second-highest-salary

SELECT (
           SELECT DISTINCT `salary`
           FROM `employee`
           ORDER BY `salary` DESC
           LIMIT 1,1) `secondhighestsalary`
```

# 0177.nth-highest-salary 找第N名 Null 也返回 
```sql
# Link: https://leetcode-cn.com/problems/nth-highest-salary

CREATE FUNCTION getNthHighestSalary(`n` INT) RETURNS INT
BEGIN
    SET `n` := `n` - 1;
    RETURN (
        SELECT (
                   SELECT DISTINCT `salary`
                   FROM `employee`
                   ORDER BY `salary` DESC
                   LIMIT `n`,1) `nthhighestsalary`
    );
END
```

# 0178.rank-scores 稠密排行 
```sql
# Link: https://leetcode-cn.com/problems/rank-scores

SELECT `score`,
       @`rank` := @`rank` + (@`pre` != (@`pre` := `score`)) `rank`
FROM `scores`,
     (SELECT @`pre` := -1, @`rank` := 0) `tmp`
ORDER BY `score` DESC;

#

SELECT `score`,
       DENSE_RANK() OVER (ORDER BY `score` DESC) `rank`
FROM `scores`
```

# 0180.consecutive-numbers 连续出现3次的数字 
```sql
# Link: https://leetcode-cn.com/problems/consecutive-numbers

SELECT DISTINCT `num` `consecutivenums`
FROM (SELECT IF(@`pre` = `num`, @`count` := @`count` + 1, @`count` := 1) `counter`,
             @`pre` := `num`                                             `num`
      FROM `logs`,
           (SELECT @`pre` := 0) AS `t`) AS `t`
WHERE 3 <= `counter`;

#

SELECT DISTINCT `num` `consecutivenums`
FROM (
         SELECT *,
                ROW_NUMBER() OVER (PARTITION BY `num` ORDER BY `id`) `rownum`
         FROM `logs`
     ) `t`
GROUP BY (`id` + 1 - `rownum`), `num`
HAVING 3 <= COUNT(*);

#

SELECT DISTINCT `l1`.`num` `consecutivenums`
FROM `logs` `l1`,
     `logs` `l2`,
     `logs` `l3`
WHERE `l1`.`id` = `l2`.`id` - 1
  AND `l2`.`id` = `l3`.`id` - 1
  AND `l1`.`num` = `l2`.`num`
  AND `l2`.`num` = `l3`.`num`
```

# 0181.employees-earning-more-than-their-managers 超过经理收入的员工 
```sql
# Link: https://leetcode-cn.com/problems/employees-earning-more-than-their-managers

SELECT `e1`.`name` `employee`
FROM `employee`               AS `e1`
         LEFT JOIN `employee` AS `e2` ON `e1`.`managerid` = `e2`.`id`
WHERE `e2`.`salary` < `e1`.`salary`
```

# 0182.duplicate-emails 找重复邮箱 
```sql
# Link: https://leetcode-cn.com/problems/duplicate-emails

SELECT `email`
FROM `person`
GROUP BY `email`
HAVING 1 < COUNT(`email`)
```

# 0183.customers-who-never-order 从不订购的客户 
```sql
# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT `customers`.`name` `customers`
FROM `customers`
         LEFT JOIN `orders` ON `customers`.`id` = `orders`.`customerid`
WHERE `orders`.`customerid` IS NULL;

#

SELECT `name` `customers`
FROM `customers`
WHERE `id` NOT IN (SELECT `customerid` FROM `orders`)
```

# 0184.department-highest-salary 部门最高工资的员工 
```sql
# Link: https://leetcode-cn.com/problems/department-highest-salary

SELECT `d`.`name` 'department',
       `e`.`name` 'employee',
       `salary`
FROM `employee`   `e`
         JOIN
     `department` `d` ON `e`.`departmentid` = `d`.`id`
         AND (`e`.`departmentid`, `salary`) IN
             (SELECT `departmentid`,
                     MAX(`salary`)
              FROM `employee`
              GROUP BY `departmentid`
             );

#

SELECT `department`,
       `employee`,
       `salary`
FROM (SELECT `d`.`name`                                                            'department',
             `e`.`name`                                                            'employee',
             `e`.`salary`,
             RANK() OVER (PARTITION BY `e`.`departmentid` ORDER BY `salary` DESC ) `r`
      FROM `department`        `d`
               JOIN `employee` `e`
                    ON `d`.`id` = `e`.`departmentid`
     ) `t`
WHERE `r` = 1;
```

# 0185.department-top-three-salaries 部门工资前三高的员工 
```sql
# Link: https://leetcode-cn.com/problems/department-top-three-salaries

SELECT `department`, `employee`, `salary`
FROM (SELECT `d`.`name`                                                                       `department`,
             `e1`.`name`                                                                      `employee`,
             `e1`.`salary`                                                                    `salary`,
             DENSE_RANK() OVER (PARTITION BY `e1`.`departmentid` ORDER BY `e1`.`salary` DESC) `r`
      FROM `employee`            `e1`
               JOIN `department` `d`
                    ON
                        `e1`.`departmentid` = `d`.`id`) `t`
WHERE `r` <= 3;

#

SELECT `d`.`name`    `department`,
       `e1`.`name`   `employee`,
       `e1`.`salary` `salary`
FROM `employee`            `e1`
         JOIN `department` `d`
              ON
                      `e1`.`departmentid` = `d`.`id`
                      AND (SELECT COUNT(DISTINCT `e2`.`salary`)
                           FROM `employee` AS `e2`
                           WHERE `e1`.`salary` < `e2`.`salary`
                             AND `e1`.`departmentid` = `e2`.`departmentid`) < 3
ORDER BY `salary` DESC;
```

# 0196.delete-duplicate-emails 删除重复邮箱 
```sql
# Link: https://leetcode-cn.com/problems/delete-duplicate-emails

DELETE `p2`
FROM `person`          `p1`
         JOIN `person` `p2`
WHERE `p1`.`email` = `p2`.`email`
  AND `p1`.`id` < `p2`.`id`
```

# 0197.rising-temperature 温度相比昨天是上升的 
```sql
# Link: https://leetcode-cn.com/problems/rising-temperature

SELECT `w2`.`id` `id`
FROM `weather`          `w2`
         JOIN `weather` `w1`
              ON DATEDIFF(`w2`.`recorddate`, `w1`.`recorddate`) = 1
                  AND `w1`.`temperature` < `w2`.`temperature`
```

# 0262.trips-and-users 非禁止用户取消率 
```sql
# Link: https://leetcode-cn.com/problems/trips-and-users

SELECT `t`.`request_at` `day`,
       ROUND(
               SUM(IF(`t`.`status` = 'completed', 0, 1)) / COUNT(`t`.`status`),
               2)       `cancellation rate`
FROM `trips` `t`
         JOIN
     `users` `u1`
         JOIN
     `users` `u2`
     ON `t`.`client_id` = `u1`.`users_id`
         AND `u1`.`banned` = 'No'
         AND `t`.`driver_id` = `u2`.`users_id`
         AND `u2`.`banned` = 'No'
         AND `t`.`request_at` BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY `t`.`request_at`
```

# 0511.game-play-analysis-i 首次登陆的时间 
```sql
# Link: https://leetcode-cn.com/problems/game-play-analysis-i

SELECT `player_id`,
       MIN(`event_date`) `first_login`
FROM `activity`
GROUP BY `player_id`
```

# 0512.game-play-analysis-ii 首次登陆的设备名称 
```sql
# Link: https://leetcode-cn.com/problems/game-play-analysis-ii

SELECT `player_id`, `device_id`
FROM `activity`
WHERE (`player_id`, `event_date`) IN (SELECT `player_id`, MIN(`event_date`)
                                      FROM `activity`
                                      GROUP BY `player_id`);

#

SELECT `player_id`, `device_id`
FROM (SELECT `player_id`,
             `device_id`,
             `event_date`,
             MIN(`event_date`) OVER (PARTITION BY `player_id` ) `m`
      FROM `activity`) `t`
WHERE `m` = `event_date`

```

# 0534.game-play-analysis-iii 每人每天累积多少时长 
```sql
# Link: https://leetcode-cn.com/problems/game-play-analysis-iii

SELECT `a1`.`player_id`         `player_id`,
       `a1`.`event_date`        `event_date`,
       SUM(`a2`.`games_played`) `games_played_so_far`
FROM `activity` `a2`
         JOIN
     `activity` `a1`
     ON `a1`.`player_id` = `a2`.`player_id`
         AND `a2`.`event_date` <= `a1`.`event_date`
GROUP BY `a1`.`player_id`, `a1`.`event_date`
```

# 0550.game-play-analysis-iv 首日后隔天登录玩家的比率 
```sql
# Link: https://leetcode-cn.com/problems/game-play-analysis-iv

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
```

# 0569.median-employee-salary 每个公司薪酬中位数 
```sql
# Link: https://leetcode-cn.com/problems/median-employee-salary

SELECT `id`, `company`, `salary`
FROM (SELECT `id`,
             `company`,
             `salary`,
             ROW_NUMBER() OVER (PARTITION BY `company` ORDER BY `salary`) `rank`,
             COUNT(*) OVER (PARTITION BY `company`)                       `count`
      FROM `employee`) `t`
WHERE `rank` BETWEEN `count` / 2 AND `count` / 2 + 1
```

# 0570.managers-with-at-least-5-direct-reports 至少5名下属的经理 
```sql
# Link: https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports

SELECT `name`
FROM `employee`                                                                              `t1`
         JOIN (SELECT `managerid` FROM `employee` GROUP BY `managerid` HAVING COUNT(*) >= 5) `t2`
              ON `t1`.`id` = `t2`.`managerid`;

#

SELECT `name`
FROM `employee`
WHERE `id` IN (SELECT DISTINCT `managerid`
               FROM (SELECT `managerid`,
                            COUNT(`managerid`) OVER (PARTITION BY `managerid`) `c`
                     FROM `employee`
                     ORDER BY `c` DESC
                    ) `t`
               WHERE `c` >= 5)
```

# 0571.find-median-given-frequency-of-numbers 频率数字的中位数 
```sql
# Link: https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers


SELECT AVG(`number`) `median`
FROM (SELECT *,
             SUM(`frequency`) OVER (ORDER BY `number` ASC)  `n1`,
             SUM(`frequency`) OVER (ORDER BY `number` DESC) `n2`
      FROM `numbers`) `t`
WHERE `n1` BETWEEN `n2` - `frequency` AND `n2` + `frequency`;

```

# 0574.winning-candidate 当选者 
```sql
# Link: https://leetcode-cn.com/problems/winning-candidate

SELECT `name`
FROM (
         SELECT `candidateid`
         FROM `vote`
         GROUP BY `candidateid`
         ORDER BY COUNT(`candidateid`) DESC
         LIMIT 1) `t`
         JOIN `candidate` ON `candidate`.`id` = `candidateid`
```

# 0577.employee-bonus 员工奖金 
```sql
# Link: https://leetcode-cn.com/problems/employee-bonus

SELECT `name`, `bonus`
FROM `employee`
         LEFT JOIN `bonus` ON `employee`.`empid` = `bonus`.`empid`
WHERE `bonus`.`bonus` < 1000
   OR `bonus` IS NULL
```

# 0578.get-highest-answer-rate-question 回答率最高的问题 
```sql
# Link: https://leetcode-cn.com/problems/get-highest-answer-rate-question

SELECT `question_id` `survey_log`
FROM `survey_log`
GROUP BY `question_id`
ORDER BY SUM(IF(`action` = 'answer', 1, 0)) / SUM(IF(`action` = 'show', 1, 0)) DESC
LIMIT 1
```

# 0579.find-cumulative-salary-of-an-employee 员工累积薪水 
```sql
# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT `e1`.`id`,
       `e1`.`month`,
       (IFNULL(`e1`.`salary`, 0) + IFNULL(`e2`.`salary`, 0) + IFNULL(`e3`.`salary`, 0)) `salary`
FROM (SELECT `id`,
             MAX(`month`) `month`
      FROM `employee`
      GROUP BY `id`
      HAVING COUNT(*) > 1) AS `maxmonth`
         LEFT JOIN
     `employee`               `e1` ON `maxmonth`.`id` = `e1`.`id`
         AND `maxmonth`.`month` > `e1`.`month`
         LEFT JOIN
     `employee`               `e2` ON `e2`.`id` = `e1`.`id`
         AND `e2`.`month` = `e1`.`month` - 1
         LEFT JOIN
     `employee`               `e3` ON `e3`.`id` = `e1`.`id`
         AND `e3`.`month` = `e1`.`month` - 2
ORDER BY `id`, `month` DESC;

#

SELECT `id`, `month`, SUM(`salary`) `salary`
FROM (
         SELECT `e1`.`id`, `e1`.`month`, `e1`.`salary`
         FROM (SELECT `id`, `month`, `salary`, MAX(`month`) OVER (PARTITION BY `id`,`month`) `maxmonth`
               FROM `employee`)   `e1`
                  JOIN `employee` `e2`
         WHERE `e1`.`month` != `e1`.`maxmonth`
           AND `e1`.`id` = `e2`.`id`
           AND `e2`.`month` BETWEEN `e1`.`month` - 2 AND `e1`.`month`
         ORDER BY `id`
     ) `t`
GROUP BY `id`, `month`
ORDER BY `id`, `month` DESC;

```

# 0580.count-student-number-in-departments 各专业学生人数 
```sql
# Link: https://leetcode-cn.com/problems/count-student-number-in-departments

SELECT `dept_name`, COUNT(`student_id`) `student_number`
FROM `department`            `d`
         LEFT JOIN `student` `s` ON `d`.`dept_id` = `s`.`dept_id`
GROUP BY `dept_name`
ORDER BY `student_number` DESC
```

# 0584.find-customer-referee 用户的推荐人 
```sql
# Link: https://leetcode-cn.com/problems/find-customer-referee


SELECT `name`
FROM `customer`
WHERE `referee_id` != 2
   OR `referee_id` IS NULL
```

# 0585.investments-in-2016 2016年的投资 
```sql
# Link: https://leetcode-cn.com/problems/investments-in-2016

SELECT ROUND(SUM(`tiv_2016`), 2) `tiv_2016`
FROM (SELECT *,
             COUNT(*) OVER ( PARTITION BY `tiv_2015`)  `y`,
             COUNT(*) OVER ( PARTITION BY `lat`,`lon`) `p`
      FROM `insurance`) `t`
WHERE `y` > 1 && `p` = 1
```

# 0586.customer-placing-the-largest-number-of-orders 订单最多的客户 
```sql
# Link: https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders

SELECT `customer_number`
FROM `orders`
GROUP BY `customer_number`
ORDER BY COUNT(*) DESC
LIMIT 1
```

# 0595.big-countries 大的国家 
```sql
# Link: https://leetcode-cn.com/problems/big-countries

SELECT `name`, `population`, `area`
FROM `world`
WHERE `area` > 3000000
   OR `population` > 25000000
```

# 0596.classes-more-than-5-students 超过5名学生的课 
```sql
# Link: https://leetcode-cn.com/problems/classes-more-than-5-students

SELECT `class`
FROM `courses`
GROUP BY `class`
HAVING COUNT(DISTINCT `student`) >= 5
```

# 0597.friend-requests-i-overall-acceptance-rate 好友申请总体通过率 
```sql
# Link: https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate

SELECT ROUND(
               IFNULL(
                           (SELECT COUNT(*)
                            FROM (SELECT DISTINCT `requester_id`, `accepter_id` FROM `requestaccepted`) AS `a`)
                           /
                           (SELECT COUNT(*)
                            FROM (SELECT DISTINCT `sender_id`, `send_to_id` FROM `friendrequest`) AS `b`),
                           0)
           , 2) `accept_rate`;
```

# 0601.human-traffic-of-stadium 人流量 
```sql
# Link: https://leetcode-cn.com/problems/human-traffic-of-stadium

WITH `countt` AS (SELECT `id`,
                         COUNT(*) OVER (PARTITION BY `rn` ORDER BY `rn` ) `counter`
                  FROM (SELECT `id`,
                               `id` - ROW_NUMBER() OVER (ORDER BY `id`) `rn`
                        FROM `stadium`
                        WHERE `people` >= 100) `rowt`)
SELECT `s`.*
FROM `stadium` `s`
         JOIN `countt` ON
    `s`.`id` = `countt`.`id`
WHERE `countt`.`counter` > 2
ORDER BY `s`.`visit_date`;
```

# 0602.friend-requests-ii-who-has-the-most-friends 谁有最多的好友 
```sql
# Link: https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends

SELECT `id`, SUM(`n`) `num`
FROM (SELECT `accepter_id` `id`, COUNT(*) `n`
      FROM `request_accepted`
      GROUP BY `accepter_id`
      UNION ALL
      SELECT `requester_id` `id`, COUNT(*) `n`
      FROM `request_accepted`
      GROUP BY `requester_id`) `t`
GROUP BY `id`
ORDER BY `num` DESC
LIMIT 1;

```

# 0603.consecutive-available-seats 连续空余座位 
```sql
# Link: https://leetcode-cn.com/problems/consecutive-available-seats

SELECT `seat_id`
FROM (SELECT `seat_id`, COUNT(*) OVER (PARTITION BY `r`) `c`
      FROM (SELECT `seat_id`,
                   `seat_id` - ROW_NUMBER() OVER (ORDER BY `seat_id`) `r`
            FROM `cinema`
            WHERE `free` = 1) `t`) `t2`
WHERE `c` > 1
GROUP BY `seat_id`

```

# 0607.sales-person 销售员 
```sql
# Link: https://leetcode-cn.com/problems/sales-person

SELECT DISTINCT `name`
FROM `salesperson`
WHERE `sales_id` NOT IN (SELECT `sales_id`
                         FROM `orders`
                         WHERE `com_id` = (SELECT `com_id` FROM `company` WHERE `name` = 'RED'))
```

# 0608.tree-node 树节点 
```sql
# Link: https://leetcode-cn.com/problems/tree-node

SELECT DISTINCT `t1`.`id`,
                IF(`t1`.`p_id` IS NULL, 'Root',
                   IF(`t2`.`id` IS NOT NULL, 'Inner', 'Leaf')) `type`
FROM `tree`               AS `t1`
         LEFT JOIN `tree` AS `t2`
                   ON `t1`.`id` = `t2`.`p_id`
```

# 0610.triangle-judgement 判断三角形 
```sql
# Link: https://leetcode-cn.com/problems/triangle-judgement

SELECT *, IF(`x` + `y` > `z` AND `x` + `z` > `y` AND `y` + `z` > `x`, "yes", "no") `triangle`
FROM `triangle`
```

# 0612.shortest-distance-in-a-plane 平面上的最近距离 
```sql
# Link: https://leetcode-cn.com/problems/shortest-distance-in-a-plane

SELECT ROUND(SQRT(MIN((POW(`p1`.`x` - `p2`.`x`, 2) + POW(`p1`.`y` - `p2`.`y`, 2)))), 2) `shortest`
FROM `point_2d` `p1`
         JOIN
     `point_2d` `p2` ON (`p1`.`x` != `p2`.`x` AND `p1`.`y` = `p2`.`y`)
         OR (`p1`.`x` != `p2`.`x` AND `p1`.`y` != `p2`.`y`)
         OR (`p1`.`x` = `p2`.`x` AND `p1`.`y` != `p2`.`y`)

```

# 0613.shortest-distance-in-a-line 直线上的最近距离 
```sql
# Link: https://leetcode-cn.com/problems/shortest-distance-in-a-line

SELECT MIN(ABS(`p1`.`x` - `p2`.`x`)) `shortest`
FROM `point`          `p1`
         JOIN `point` `p2` ON `p1`.`x` != `p2`.`x`;

```

# 0614.second-degree-follower 二级关注者 
```sql
# Link: https://leetcode-cn.com/problems/second-degree-follower

SELECT `followee`                 `follower`,
       COUNT(DISTINCT `follower`) `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`



SELECT `followee`                 `follower`,
       COUNT(DISTINCT `follower`) `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`
```

# 0615.average-salary-departments-vs-company 平均工资：部门与公司比较 
```sql
# Link: https://leetcode-cn.com/problems/average-salary-departments-vs-company

SELECT DISTINCT DATE_FORMAT(`pay_date`, '%Y-%m')                        `pay_month`,
                `department_id`,
                IF(`b` = `a`, 'same', IF(`a` < `b`, 'lower', 'higher')) `comparison`
FROM (SELECT `department_id`,
             `amount`,
             `pay_date`,
             AVG(`amount`) OVER (PARTITION BY `pay_date`,`department_id`) `a`,
             AVG(`amount`) OVER (PARTITION BY `pay_date`)                 `b`
      FROM `salary`            `st`
               JOIN `employee` `et` ON `st`.`employee_id` = `et`.`employee_id`) `t`
ORDER BY `pay_month` DESC;
```

# 0618.students-report-by-geography 学生地理信息报告 
```sql
# Link: https://leetcode-cn.com/problems/students-report-by-geography

SELECT `america`, `asia`, `europe`
FROM (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `america`
      FROM `student`
      WHERE `continent` = 'America')              `a`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `asia`
                    FROM `student`
                    WHERE `continent` = 'Asia')   `b` ON `a`.`r` = `b`.`r`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `europe`
                    FROM `student`
                    WHERE `continent` = 'Europe') `c` ON `a`.`r` = `c`.`r`
```

# 0619.biggest-single-number 最出现一次的最大数字 
```sql
# Link: https://leetcode-cn.com/problems/biggest-single-number

SELECT (SELECT `num`
        FROM `my_numbers`
        GROUP BY `num`
        HAVING COUNT(*) = 1
        ORDER BY `num` DESC
        LIMIT 1) `num`
```

# 0620.not-boring-movies 有趣的电影 
```sql
# Link: https://leetcode-cn.com/problems/not-boring-movies

SELECT `id`, `movie`, `description`, `rating`
FROM `cinema`
WHERE `description` != 'boring'
  AND `id` % 2 = 1
ORDER BY `rating` DESC
```

# 0626.exchange-seats 换座位 
```sql
# Link: https://leetcode-cn.com/problems/exchange-seats

SELECT ROW_NUMBER() OVER (ORDER BY (`id` + 1 - 2 * POWER(0, `id` % 2))) `id`,
       `student`
FROM `seat`



SELECT IF(`id` % 2 = 0, `id` - 1, `id` + 1) `id`
    `student`
FROM `order` BY id ASC;

```

# 0627.swap-salary 变更性别 
```sql
# Link: https://leetcode-cn.com/problems/swap-salary

UPDATE `salary`
SET `sex` = IF(`sex` = 'f', 'm', 'f')
```

# 1045.customers-who-bought-all-products 买下所有产品的客户 
```sql
# Link: https://leetcode-cn.com/problems/customers-who-bought-all-products

SELECT `customer_id`
FROM `customer`
GROUP BY `customer_id`
HAVING COUNT(DISTINCT `product_key`) = (SELECT COUNT(*) `cc` FROM `product`)
```

# 1050.actors-and-directors-who-cooperated-at-least-three-times 合作至少三次的演员和导员 
```sql
# Link: https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times

SELECT `actor_id`, `director_id`
FROM `actordirector`
GROUP BY `actor_id`, `director_id`
HAVING COUNT(*) >= 3
```

# 1068.product-sales-analysis-i 产品的年份和价格 
```sql
# Link: https://leetcode-cn.com/problems/product-sales-analysis-i

SELECT `product_name`, `year`, `price`
FROM `sales`
         JOIN `product` ON `sales`.`product_id` = `product`.`product_id`
```

# 1069.product-sales-analysis-ii 产品的销售总额 
```sql
# Link: https://leetcode-cn.com/problems/product-sales-analysis-ii

SELECT `product_id`, SUM(`quantity`) `total_quantity`
FROM `sales`
GROUP BY `product_id`
```

# 1070.product-sales-analysis-iii 第一年的价格 
```sql
# Link: https://leetcode-cn.com/problems/product-sales-analysis-iii


SELECT `product_id`, `year` `first_year`, `quantity`, `price`
FROM `sales`
WHERE (`product_id`, `year`) IN (SELECT `product_id`, MIN(`year`)
                                 FROM `sales`
                                 GROUP BY `product_id`);
```

# 1075.project-employees-i 项目的平均年限 
```sql
# Title: Project Employees I
# Link: https://leetcode-cn.com/problems/project-employees-i

SELECT `project_id`, ROUND(AVG(`experience_years`), 2) `average_years`
FROM `project`                 AS `p`
         INNER JOIN `employee` AS `e`
                    ON `p`.`employee_id` = `e`.`employee_id`
GROUP BY `project_id`;
```

# 1076.project-employees-ii 员工最多的项目 销售额最高的销售者 
```sql
# Title: Project Employees II
# Link: https://leetcode-cn.com/problems/project-employees-ii

WITH `tmp` AS (SELECT `project_id`, COUNT(*) `c` FROM `project` GROUP BY `project_id`)
SELECT `project_id`
FROM `tmp`
WHERE `c` = (SELECT MAX(`c`) FROM `tmp`);
```

# 1077.project-employees-iii 项目经济最丰富的员工 
```sql
# Link: https://leetcode-cn.com/problems/project-employees-iii

SELECT `project_id`, `employee_id`
FROM (SELECT `project_id`,
             `p`.`employee_id`,
             RANK() OVER (PARTITION BY `project_id` ORDER BY `experience_years` DESC) `r`
      FROM `project`                 `p`
               INNER JOIN `employee` `e`
                          ON `p`.`employee_id` = `e`.`employee_id`) `a`
WHERE `r` = 1;
```

# 1082.sales-analysis-i 销售额最高的销售者 
```sql
# Title: Sales Analysis I
# Link: https://leetcode-cn.com/problems/sales-analysis-i

SELECT `seller_id`
FROM (
         SELECT `seller_id`, DENSE_RANK() OVER (ORDER BY `total` DESC) `n`
         FROM (
                  SELECT `seller_id`, SUM(`price`) `total` FROM `sales` GROUP BY `seller_id`) `t1`) `t2`
WHERE `n` = 1
```

# 1083.sales-analysis-ii 买 S8 却没有 iPhone 的买家 
```sql
# Link: https://leetcode-cn.com/problems/sales-analysis-ii

SELECT `buyer_id`
FROM (SELECT `buyer_id`,
             IF(`p`.`product_name` = 'S8', 1, 0)     `s8`,
             IF(`p`.`product_name` = 'iPhone', 1, 0) `ip`
      FROM `sales`            `s`
               JOIN `product` `p` ON `s`.`product_id` = `p`.`product_id`) `t`
GROUP BY `buyer_id`
HAVING 0 < SUM(`ip`)
   AND SUM(`s8`) = 0
```

# 1084.sales-analysis-iii 只在春季销售的产品 
```sql
# Link: https://leetcode-cn.com/problems/sales-analysis-iii

SELECT `p`.`product_id`, `p`.`product_name`
FROM `product`        `p`
         JOIN `sales` `s` ON `p`.`product_id` = `s`.`product_id`
GROUP BY `product_id`
HAVING (SUM(`sale_date` BETWEEN '2019-01-01' AND '2019-03-31')) = COUNT(*);

#

SELECT `product_id`,
       `product_name`
FROM `product`
WHERE `product_id` NOT IN
      (SELECT `product_id`
       FROM `sales`
       WHERE `sale_date` NOT BETWEEN '2019-01-01' AND '2019-03-31')
```

# 1097.game-play-analysis-v 第二天留存率 
```sql
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

```

# 1098.unpopular-books 过去一年少于十本的书 
```sql
# Link: https://leetcode-cn.com/problems/unpopular-books


SELECT `b`.`book_id`, `b`.`name`
FROM `books`                `b`
         LEFT JOIN `orders` `o` ON `b`.`book_id` = `o`.`book_id`
    AND '2018-06-23' <= `o`.`dispatch_date`
WHERE `b`.`available_from` < '2019-05-23'
GROUP BY `b`.`book_id`
HAVING IFNULL(SUM(`o`.`quantity`), 0) < 10
```

# 1107.new-users-daily-count 每日新用户 
```sql
# Link: https://leetcode-cn.com/problems/new-users-daily-count


SELECT `activity_date` `login_date`, COUNT(*) `user_count`
FROM (
         SELECT `user_id`, MIN(`activity_date`) `activity_date`
         FROM `traffic`
         WHERE `activity` = 'login'
         GROUP BY `user_id`) `t`
WHERE '2019-03-31' < `activity_date`
GROUP BY `login_date
```

# 1112.highest-grade-for-each-student 最高分数的学科 
```sql
# Link: https://leetcode-cn.com/problems/highest-grade-for-each-student

SELECT `student_id`, `course_id`, `grade`
FROM (SELECT *, ROW_NUMBER() OVER (PARTITION BY `student_id` ORDER BY `grade` DESC,`course_id`) `r`
      FROM `enrollments`) `t`
WHERE `r` = 1
```

# 1113.reported-posts 不同的报告记录 
```sql
# Link: https://leetcode-cn.com/problems/reported-posts

SELECT `extra` "report_reason", COUNT(DISTINCT `post_id`) "report_count"
FROM `actions`
WHERE `action_date` = '2019-07-04'
  AND `action` = 'report'
  AND `extra` IS NOT NULL
GROUP BY `extra`
```

# 1126.active-businesses 活跃的业务 
```sql
# Link: https://leetcode-cn.com/problems/active-businesses

SELECT `business_id`
FROM (SELECT *,
             AVG(`occurences`) OVER (PARTITION BY `event_type`) `a`
      FROM `events`) `t`
WHERE `occurences` > `a`
GROUP BY `business_id`
HAVING COUNT(*) > 1
```

# 1127.user-purchase-platform 统计单端或双端人数 
```sql
# Link: https://leetcode-cn.com/problems/user-purchase-platform

SELECT `spend_date`,
       `t2`.`platform`,
       SUM(IF(`t1`.`platform` = `t2`.`platform`, `amount`, 0)) `total_amount`,
       COUNT(IF(`t1`.`platform` = `t2`.`platform`, 1, NULL))   `total_users`
FROM (
         SELECT `spend_date`,
                `user_id`,
                IF(COUNT(DISTINCT `platform`) = 2, 'both', `platform`) `platform`,
                SUM(`amount`)                                          `amount`
         FROM `spending`
         GROUP BY `user_id`, `spend_date`
     ) `t1`
         JOIN
     (
         SELECT 'desktop' `platform`
         UNION
         SELECT 'mobile' `platform`
         UNION
         SELECT 'both' `platform`
     ) `t2`
GROUP BY `spend_date`, `platform`

```

# 1132.reported-posts-ii 垃圾清除率 
```sql
# Link: https://leetcode-cn.com/problems/reported-posts-ii

SELECT ROUND(AVG(`proportion`) * 100, 2) `average_daily_percent`
FROM (
         SELECT `a`.`action_date`,
                COUNT(DISTINCT `r`.`post_id`) / COUNT(DISTINCT `a`.`post_id`) `proportion`
         FROM `actions`                `a`
                  LEFT JOIN `removals` `r`
                            ON `a`.`post_id` = `r`.`post_id`
         WHERE `extra` = 'spam'
         GROUP BY `a`.`action_date`
     ) `a`;

```

# 1141.user-activity-for-the-past-30-days-i 近30天活跃数 
```sql
# Link: https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i


SELECT IFNULL(
               ROUND(
                       COUNT(DISTINCT `session_id`) / COUNT(DISTINCT `user_id`),
                       2),
               0) `average_sessions_per_user`
FROM `activity`
WHERE '2019-06-27' < `activity_date`
```

# 1142.user-activity-for-the-past-30-days-ii 平均会话次数 
```sql
# Link: https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii


```

# 1148.article-views-i 浏览过自己文章的作者 
```sql
# Link: https://leetcode-cn.com/problems/article-views-i

SELECT DISTINCT `author_id` `id`
FROM `views`
WHERE `viewer_id` = `author_id`
ORDER BY `author_id`
```

# 1149.article-views-ii 一天浏览过两篇文章的人 
```sql
# Link: https://leetcode-cn.com/problems/article-views-ii

SELECT DISTINCT `viewer_id` `id`
FROM `views`
GROUP BY `viewer_id`, `view_date`
HAVING 1 < COUNT(DISTINCT `article_id`)
ORDER BY `id`
```

# 1158.market-analysis-i 统计2019年订单总数 
```sql
# Link: https://leetcode-cn.com/problems/market-analysis-i


SELECT `u`.`user_id`     `buyer_id`,
       `join_date`,
       COUNT(`order_id`) `orders_in_2019`
FROM `users`                `u`
         LEFT JOIN `orders` `o` ON `u`.`user_id` = `o`.`buyer_id` AND `order_date` BETWEEN '2019-01-01' AND '2019-12-31'
GROUP BY `user_id`
```

# 1159.market-analysis-ii 售出第二件不是喜欢的商品 
```sql
# Link: https://leetcode-cn.com/problems/market-analysis-ii


SELECT `user_id`                                                                        `seller_id`,
       IF(`t`.`item_brand` != `u`.`favorite_brand` || `seller_id` IS NULL, 'no', 'yes') `2nd_item_fav_brand`
FROM `users`                                                        `u`
         LEFT JOIN (SELECT `seller_id`,
                           `item_brand`,
                           RANK() OVER (PARTITION BY `seller_id` ORDER BY `order_date`) `r`
                    FROM `orders`         AS `o`
                             JOIN `items` AS `i`
                                  ON `o`.`item_id` = `i`.`item_id`) `t`
                   ON `u`.`user_id` = `t`.`seller_id`
                       AND `r` = 2;
```

# 1164.product-price-at-a-given-date 变更后的价格 
```sql
# Link: https://leetcode-cn.com/problems/product-price-at-a-given-date

SELECT DISTINCT `product_id`, `price`
FROM (
         (SELECT `product_id`, @`price` := 10 `price`
          FROM `products` `p`
          GROUP BY `product_id`)
         UNION
         SELECT `product_id`, `price`
         FROM (SELECT `product_id`,
                      RANK() OVER (PARTITION BY `product_id` ORDER BY `change_date` DESC) `r`,
                      `new_price`                                                         `price`
               FROM `products`
               WHERE `change_date` <= '2019-08-16'
               GROUP BY `product_id`) `t2`
         WHERE `r` = 1
     ) `t`;

SELECT DISTINCT `p`.`product_id`, IFNULL(`t`.`new_price`, 10) `price`
FROM `products`                                          `p`
         LEFT JOIN (SELECT `product_id`,
                           `new_price`,
                           RANK() OVER (PARTITION BY `product_id` ORDER BY `change_date` DESC) `r`
                    FROM `products`
                    WHERE `change_date` <= '2019-08-16') `t`
                   ON `p`.`product_id` = `t`.`product_id`
                       AND `r` = 1;

```

# 1173.immediate-food-delivery-i 当天配送率 
```sql
# Link: https://leetcode-cn.com/problems/immediate-food-delivery-i

SELECT ROUND(
                       SUM(IF(`order_date` = `customer_pref_delivery_date`, 1, 0)) /
                       COUNT(*) * 100, 2) `immediate_percentage`
FROM `delivery`
```

# 1174.immediate-food-delivery-ii 首日即时配送订单 
```sql
# Link: https://leetcode-cn.com/problems/immediate-food-delivery-ii

SELECT ROUND(
                   SUM(`order_date` = `customer_pref_delivery_date`) * 100 /
                   COUNT(*),
                   2
           ) `immediate_percentage`
FROM `delivery`
WHERE (`customer_id`, `order_date`) IN (
    SELECT `customer_id`, MIN(`order_date`)
    FROM `delivery`
    GROUP BY `customer_id`
)
```

# 1179.reformat-department-table 格式化工资 
```sql
# Link: https://leetcode-cn.com/problems/reformat-department-table

SELECT `id`,
       SUM(IF(`month` = 'Jan', `revenue`, NULL)) `jan_revenue`,
       SUM(IF(`month` = 'Feb', `revenue`, NULL)) `feb_revenue`,
       SUM(IF(`month` = 'Mar', `revenue`, NULL)) `mar_revenue`,
       SUM(IF(`month` = 'Apr', `revenue`, NULL)) `apr_revenue`,
       SUM(IF(`month` = 'May', `revenue`, NULL)) `may_revenue`,
       SUM(IF(`month` = 'Jun', `revenue`, NULL)) `jun_revenue`,
       SUM(IF(`month` = 'Jul', `revenue`, NULL)) `jul_revenue`,
       SUM(IF(`month` = 'Aug', `revenue`, NULL)) `aug_revenue`,
       SUM(IF(`month` = 'Sep', `revenue`, NULL)) `sep_revenue`,
       SUM(IF(`month` = 'Oct', `revenue`, NULL)) `oct_revenue`,
       SUM(IF(`month` = 'Nov', `revenue`, NULL)) `nov_revenue`,
       SUM(IF(`month` = 'Dec', `revenue`, NULL)) `dec_revenue`
FROM `department`
GROUP BY `id`
ORDER BY `id`;
```

# 1193.monthly-transactions-i 统计 
```sql
# Link: https://leetcode-cn.com/problems/monthly-transactions-i


SELECT DATE_FORMAT(`trans_date`, '%Y-%m')         `month`,
       `country`,
       COUNT(*)                                   `trans_count`,
       COUNT(IF(`state` = 'approved', 1, NULL))   `approved_count`,
       SUM(`amount`)                              `trans_total_amount`,
       SUM(IF(`state` = 'approved', `amount`, 0)) `approved_total_amount`
FROM `transactions`
GROUP BY `month`, `country`
```

# 1194.tournament-winners 每组的优胜者 
```sql
# Link: https://leetcode-cn.com/problems/tournament-winners

SELECT `group_id`, `player_id`
FROM (
         SELECT `group_id`, `player_id`, SUM(`score`) `score`
         FROM (
                  -- 每个用户总的 first_score
                  SELECT `p1`.`group_id`, p1.`player_id`, SUM(`m1`.`first_score`) `score`
                  FROM `players`          `p1`
                           JOIN `matches` m1 ON `p1`.`player_id` = `m1`.`first_player`
                  GROUP BY `p1`.`player_id`

                  UNION ALL

                  -- 每个用户总的 second_score
                  SELECT `p2`.`group_id`, `p2`.`player_id`, SUM(`m2`.`second_score`) `score`
                  FROM `players`          `p2`
                           JOIN `matches` m2 ON `p2`.`player_id` = `m2`.`second_player`
                  GROUP BY `p2`.`player_id`
              ) `s`
         GROUP BY `player_id`
         ORDER BY `score` DESC, `player_id`
     ) `t`
GROUP BY `group_id`
ORDER BY `group_id`
```

# 1204.last-person-to-fit-in-the-elevator 最后进入电梯 
```sql
# Link: https://leetcode-cn.com/problems/last-person-to-fit-in-the-elevator


SELECT `a`.`person_name`
FROM (
         SELECT `person_name`, @`pre` := @`pre` + `weight` `weight`
         FROM `queue`,
              (SELECT @`pre` := 0) `tmp`
         ORDER BY `turn`
     ) `a`
WHERE `a`.`weight` <= 1000
ORDER BY `a`.`weight` DESC
LIMIT 1
```

# 1205.monthly-transactions-ii 交易统计 
```sql
# Link: https://leetcode-cn.com/problems/monthly-transactions-ii

SELECT LEFT(`trans_date`, 7)                        `month`,
       `country`,
       SUM(IF(`state` = 'approved', 1, 0))          `approved_count`,
       SUM(IF(`state` = 'approved', `amount`, 0))   `approved_amount`,
       SUM(IF(`state` = 'chargeback', 1, 0))        `chargeback_count`,
       SUM(IF(`state` = 'chargeback', `amount`, 0)) `chargeback_amount`
FROM (SELECT `id`, `country`, 'chargeback' `state`, `amount`, `c`.`trans_date`
      FROM `transactions`         `t1`
               JOIN `chargebacks` `c`
                    ON `t1`.`id` = `c`.`trans_id`
      UNION
      SELECT *
      FROM `transactions`) `t2`
GROUP BY `month`, `country`
HAVING SUM(IF(`state` = 'approved' OR `state` = 'chargeback', 1, 0)) > 0
ORDER BY `month`;
```

# 1211.queries-quality-and-percentage 
```sql
# Link: https://leetcode-cn.com/problems/queries-quality-and-percentage

SELECT `query_name`,
       ROUND(AVG(`rating` / `position`), 2)                   `quality`,
       ROUND(SUM(IF(`rating` < 3, 1, 0)) * 100 / COUNT(*), 2) `poor_query_percentage`
FROM `queries`
GROUP BY `query_name`
```

# 1212.team-scores-in-football-tournament 计算得分 
```sql
# Link: https://leetcode-cn.com/problems/team-scores-in-football-tournament


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

```

# 1225.report-contiguous-dates 统计连续 
```sql
# Link: https://leetcode-cn.com/problems/report-contiguous-dates


SELECT `type` `period_state`, MIN(`date`) `start_date`, MAX(`date`) `end_date`
FROM (
         SELECT `type`, `date`, SUBDATE(`date`, ROW_NUMBER() OVER (PARTITION BY `type` ORDER BY `date`)) `diff`
         FROM (
                  SELECT 'failed' `type`, `fail_date` `date`
                  FROM `failed`
                  UNION ALL
                  SELECT 'succeeded' `type`, `success_date` `date`
                  FROM `succeeded`
              ) `t1`
     ) `t2`
WHERE `date` BETWEEN '2019-01-01' AND '2019-12-31'
GROUP BY `type`, `diff`
ORDER BY `start_date`
```

# 1241.number-of-comments-per-post 查询评论数 
```sql
# Link: https://leetcode-cn.com/problems/number-of-comments-per-post

SELECT `post_id`, COUNT(`sub_id`) `number_of_comments`
FROM (
         SELECT DISTINCT `post`.`sub_id` `post_id`, `sub`.`sub_id` `sub_id`
         FROM `submissions`               `post`
                  LEFT JOIN `submissions` `sub`
                            ON `post`.`sub_id` = `sub`.`parent_id`
         WHERE `post`.`parent_id` IS NULL
     ) `t`
GROUP BY `post_id`
ORDER BY `post_id`

```

# 1251.average-selling-price 
```sql
# Link: https://leetcode-cn.com/problems/average-selling-price

SELECT `product_id`,
       ROUND(SUM(`sales`) / SUM(`units`), 2) `average_price`
FROM (
         SELECT `p`.`product_id`          `product_id`,
                `p`.`price` * `u`.`units` `sales`,
                `u`.`units`               `units`
         FROM `prices`             `p`
                  JOIN `unitssold` `u` ON `p`.`product_id` = `u`.`product_id`
         WHERE `u`.`purchase_date` BETWEEN `p`.`start_date` AND `p`.`end_date`
     ) `t`
GROUP BY `product_id`
```

# 1264.page-recommendations 推荐朋友喜欢的页面 
```sql
# Link: https://leetcode-cn.com/problems/page-recommendations

SELECT DISTINCT `page_id` `recommended_page`
FROM `likes`
WHERE `user_id` IN (
    SELECT (
               CASE
                   WHEN `user1_id` = 1 THEN `user2_id`
                   WHEN `user2_id` = 1 THEN `user1_id`
                   END
               ) `user_id`
    FROM `friendship`
    WHERE `user1_id` = 1
       OR `user2_id` = 1
)
  AND `page_id` NOT IN (
    SELECT `page_id`
    FROM `likes`
    WHERE `user_id` = 1
)
```

# 1270.all-people-report-to-the-given-manager 递归查询 
```sql
# Link: https://leetcode-cn.com/problems/all-people-report-to-the-given-manager

SELECT `e1`.`employee_id`
FROM `employees`          `e1`
         JOIN `employees` `e2` ON `e1`.`manager_id` = `e2`.`employee_id`
         JOIN `employees` `e3` ON `e2`.`manager_id` = `e3`.`employee_id`
WHERE `e1`.`employee_id` != 1
  AND `e3`.`manager_id` = 1

```

# 1280.students-and-examinations 学生各科测试次数 
```sql
# Link: https://leetcode-cn.com/problems/students-and-examinations

SELECT `s1`.`student_id`, `s1`.`student_name`, `s2`.`subject_name`, COUNT(`e`.`subject_name`) `attended_exams`
FROM `students`                   `s1`
         JOIN      `subjects`     `s2`
         LEFT JOIN `examinations` `e`
                   ON `s1`.`student_id` = `e`.`student_id`
                       AND `s2`.`subject_name` = `e`.`subject_name`
GROUP BY `s1`.`student_id`, `s1`.`student_name`, `s2`.`subject_name`
ORDER BY `s1`.`student_id`, `s2`.`subject_name`;
```

# 1285.find-the-start-and-end-number-of-continuous-ranges 区间次数 
```sql
# Link: https://leetcode-cn.com/problems/find-the-start-and-end-number-of-continuous-ranges

SELECT MIN(`log_id`) `start_id`,
       MAX(`log_id`) `end_id`
FROM (
         SELECT DISTINCT `log_id`,
                         `log_id` - ROW_NUMBER() OVER ( ORDER BY `log_id` ASC ) `r`
         FROM `logs`
     ) `t`
GROUP BY `r`
ORDER BY `start_id`

```

# 1294.weather-type-in-each-country 11月份的天气 
```sql
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
```

# 1303.find-the-team-size 团队人数 
```sql
# Link: https://leetcode-cn.com/problems/find-the-team-size

SELECT `employee_id`, COUNT(*) OVER (PARTITION BY `team_id`) `team_size`
FROM `employee`

```

# 1308.running-total-for-different-genders 男女的累积分数 
```sql
# Link: https://leetcode-cn.com/problems/running-total-for-different-genders

SELECT `gender`,
       `day`,
       SUM(`score_points`) OVER (PARTITION BY `gender` ORDER BY `day`) `total`
FROM `scores`;
```

# 1321.restaurant-growth 最近七天平均值 
```sql
# Link: https://leetcode-cn.com/problems/restaurant-growth

# Write your MySQL query statement below
SELECT `visited_on`,
       SUM(`amount`) OVER (ORDER BY `visited_on` ROWS 6 PRECEDING )               `amount`,
       ROUND(SUM(`amount`) OVER (ORDER BY `visited_on` ROWS 6 PRECEDING ) / 7, 2) `average_amount`
FROM (
         SELECT `visited_on`,
                SUM(`amount`) `amount`
         FROM `customer`
         GROUP BY `visited_on`) `a`
ORDER BY `visited_on`
LIMIT 6, 18446744073709551615


```

# 1322.ads-performance 广告效果统计 
```sql
# Link: https://leetcode-cn.com/problems/ads-performance

SELECT `ad_id`,
       ROUND(IFNULL(SUM(`action` = 'Clicked') /
                    (SUM(`action` = 'Clicked') + SUM(`action` = 'Viewed')) * 100, 0), 2) `ctr`
FROM `ads`
GROUP BY `ad_id`
ORDER BY `ctr` DESC, `ad_id`
```

# 1327.list-the-products-ordered-in-a-period 指定日期总数 
```sql
# Link: https://leetcode-cn.com/problems/list-the-products-ordered-in-a-period

SELECT `product_name`, SUM(`unit`) `unit`
FROM `products`        `p`
         JOIN `orders` `o`
              ON `p`.`product_id` = `o`.`product_id`
WHERE LEFT(`order_date`, 7) = '2020-02'
GROUP BY `product_name`
HAVING SUM(`unit`) >= 100;
```

# 1336.number-of-transactions-per-visit 每次访问的交易次数 
```sql
# Link: https://leetcode-cn.com/problems/number-of-transactions-per-visit

WITH `cte` AS (SELECT COUNT(`amount`) `transactions_count`
               FROM `visits`                     `v`
                        LEFT JOIN `transactions` `t2`
                                  ON `v`.`user_id` = `t2`.`user_id`
                                      AND `v`.`visit_date` = `t2`.`transaction_date`
               GROUP BY `v`.`user_id`, `v`.`visit_date`)

SELECT `t1`.`transactions_count`, COUNT(`cte`.`transactions_count`) `visits_count`
FROM (SELECT 0 `transactions_count`
      UNION
      SELECT ROW_NUMBER() OVER () `transactions_count`
      FROM `transactions`) `t1`
         LEFT JOIN `cte`
                   ON `t1`.`transactions_count` = `cte`.`transactions_count`
GROUP BY 1
HAVING `t1`.`transactions_count` <= (SELECT MAX(`transactions_count`)
                                     FROM `cte`);

```

# 1341.movie-rating 评论数和评分数统计 
```sql
# Link: https://leetcode-cn.com/problems/movie-rating

(SELECT `name` `results`
 FROM `users`                 `u`
          JOIN `movie_rating` `mr`
               ON `u`.`user_id` = `mr`.`user_id`
 GROUP BY `u`.`user_id`
 ORDER BY COUNT(*) DESC, `name`
 LIMIT 1)
UNION
(SELECT `title` `results`
 FROM `movies`                `m`
          JOIN `movie_rating` `mr`
               ON `m`.`movie_id` = `mr`.`movie_id`
 WHERE LEFT(`created_at`, 7) = '2020-02'
 GROUP BY `m`.`movie_id`
 ORDER BY AVG(`rating`) DESC, `title`
 LIMIT 1);
```

# 1350.students-with-invalid-departments 不存在院校的学生 
```sql
# Link: https://leetcode-cn.com/problems/students-with-invalid-departments

SELECT `id`, `name`
FROM `students`
WHERE `department_id` NOT IN (SELECT `id` FROM `departments`)
```

# 1355.activity-participants 去除第一和倒数第一 
```sql
# Link: https://leetcode-cn.com/problems/activity-participants


SELECT `activity`
FROM (SELECT `activity`,
             DENSE_RANK() OVER (ORDER BY COUNT(*))      `r1`,
             DENSE_RANK() OVER (ORDER BY COUNT(*) DESC) `r2`
      FROM `friends`
      GROUP BY `activity`) `t`
WHERE `r1` != 1
  AND `r2` != 1
```

# 1364.number-of-trusted-contacts-of-a-customer 互为好友的顾客 
```sql
# Link: https://leetcode-cn.com/problems/number-of-trusted-contacts-of-a-customer

SELECT `invoice_id`,
       `cu1`.`customer_name`,
       `price`,
       COUNT(`co`.`contact_name`)   `contacts_cnt`,
       COUNT(`cu2`.`customer_name`) `trusted_contacts_cnt`
FROM `invoices`                `i`
         JOIN      `customers` `cu1` ON `i`.`user_id` = `cu1`.`customer_id`
         LEFT JOIN `contacts`  `co` ON `i`.`user_id` = `co`.`user_id`
         LEFT JOIN `customers` `cu2` ON `co`.`contact_email` = `cu2`.`email`
GROUP BY `i`.`invoice_id`
ORDER BY `i`.`invoice_id`
```

# 1369.get-the-second-most-recent-activity 倒数第二次活动 
```sql
# Link: https://leetcode-cn.com/problems/get-the-second-most-recent-activity

SELECT `username`, `activity`, `startdate`, `enddate`
FROM (SELECT *,
             RANK() OVER (PARTITION BY `username` ORDER BY `startdate` DESC) `r`,
             COUNT(*) OVER (PARTITION BY `username`)                         `count`
      FROM `useractivity`) `t`
WHERE `r` = 2
   OR `count` = 1;


```

# 1378.replace-employee-id-with-the-unique-identifier 使用uuid替换id 
```sql
# Link: https://leetcode-cn.com/problems/replace-employee-id-with-the-unique-identifier

SELECT `unique_id`, `name`
FROM `employeeuni`
         RIGHT JOIN `employees` ON `employees`.`id` = `employeeuni`.`id`
```

# 1384.total-sales-amount-by-year 年度统计 
```sql
# Link: https://leetcode-cn.com/problems/total-sales-amount-by-year


SELECT `sales`.`product_id`,
       `product_name`,
       '2018' `report_year`,
       IF(`period_start` < '2019-01-01',
          (DATEDIFF(IF(`period_end` < '2019-01-01', `period_end`, DATE('2018-12-31')),
                    IF(`period_start` >= '2018-01-01', `period_start`, DATE('2018-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON `sales`.`product_id` = `product`.`product_id`
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2019' `report_year`,
       IF(`period_start` < '2020-01-01',
          (DATEDIFF(IF(`period_end` < '2020-01-01', `period_end`, DATE('2019-12-31')),
                    IF(`period_start` >= '2019-01-01', `period_start`, DATE('2019-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2020' `report_year`,
       (DATEDIFF(IF(`period_end` < '2021-01-01', `period_end`, DATE('2020-12-31')),
                 IF(`period_start` >= '2020-01-01', `period_start`, DATE('2020-01-01'))) + 1) *
       `average_daily_sales` `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

ORDER BY `product_id`, `report_year`



```

# 1393.capital-gainloss  资本损益 
```sql
# Link: https://leetcode-cn.com/problems/capital-gainloss

SELECT `stock_name`,
       SUM(IF(`operation` = 'Buy', -`price`, `price`)) `capital_gain_loss`
FROM `stocks`
GROUP BY `stock_name`;
```

# 1398.customers-who-bought-products-a-and-b-but-not-c 买AB不买C 
```sql
# Link: https://leetcode-cn.com/problems/customers-who-bought-products-a-and-b-but-not-c

SELECT `c`.`customer_id`,
       `c`.`customer_name`
FROM `orders`                  `o`
         LEFT JOIN `customers` `c` ON `o`.`customer_id` = `c`.`customer_id`
GROUP BY `c`.`customer_id`
HAVING SUM(`product_name` = 'A') * SUM(`product_name` = 'B') > 0
   AND SUM(`product_name` = 'C') = 0
ORDER BY `c`.`customer_id`

```

# 1407.top-travellers 旅行距离排名 
```sql
# Link: https://leetcode-cn.com/problems/top-travellers

SELECT `name`, IFNULL(SUM(`distance`), 0) `travelled_distance`
FROM `users`
         LEFT JOIN `rides` ON `users`.`id` = `rides`.`user_id`
GROUP BY `name`
ORDER BY `travelled_distance` DESC, `name`;

```

# 1412.find-the-quiet-students-in-all-exams 成绩中游学生 
```sql
# Link: https://leetcode-cn.com/problems/find-the-quiet-students-in-all-exams

SELECT `s`.`student_id`,
       `s`.`student_name`
FROM `student` `s`
WHERE `s`.`student_id` IN
      (
          SELECT `student_id`
          FROM (
                   SELECT `student_id`,
                          `score`,
                          `exam_id`,
                          MAX(`score`) OVER (PARTITION BY `exam_id`) `max_score`,
                          MIN(`score`) OVER (PARTITION BY `exam_id`) `min_score`
                   FROM `exam`
               ) `e`
          GROUP BY `student_id`
          HAVING SUM(IF(`min_score` < `e`.`score` AND `e`.`score` < `max_score`, 1, 0))
                     = COUNT(DISTINCT `exam_id`)
      )
ORDER BY `s`.`student_id`

```

# 1418.display-table-of-food-orders-in-a-restaurant 点菜展示表 
```sql
package main

import (
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant

func displayTable(orders [][]string) [][]string {
	nameSet := map[string]struct{}{}
	foodsCnt := map[int]map[string]int{}
	for _, order := range orders {
		id, _ := strconv.Atoi(order[1])
		name := order[2]
		nameSet[name] = struct{}{}

		if foodsCnt[id] == nil {
			foodsCnt[id] = map[string]int{}
		}
		foodsCnt[id][name] += 1
	}

	var names []string
	for name := range nameSet {
		names = append(names, name)
	}
	sort.Strings(names)

	var ids []int
	for id := range foodsCnt {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	l1 := len(names)
	l2 := len(ids)
	table := make([][]string, l2+1)
	table[0] = []string{"Table"}
	table[0] = append(table[0], names...)
	for i, id := range ids {
		table[i+1] = make([]string, l1+1)
		table[i+1][0] = strconv.Itoa(id)
		for j, name := range names {
			table[i+1][j+1] = strconv.Itoa(foodsCnt[id][name])
		}
	}
	return table
}

```

# 1421.npv-queries 连表 
```sql
# Link: https://leetcode-cn.com/problems/npv-queries

SELECT `q`.`id`,
       `q`.`year`,
       IFNULL(`npv`, 0) `npv`
FROM `queries` `q`
         LEFT JOIN
     `npv`     `n` USING (`id`, `year`)

```

# 1435.create-a-session-bar-chart 分段汇总 
```sql
# Link: https://leetcode-cn.com/problems/create-a-session-bar-chart


SELECT `a`.`bin`, COUNT(`b`.`bin`) `total`
FROM (
         SELECT '[0-5>' `bin` UNION SELECT '[5-10>' `bin` UNION SELECT '[10-15>' `bin` UNION SELECT '15 or more' `bin`
     ) `a`
         LEFT JOIN
     (
         SELECT CASE
                    WHEN `duration` < 300 THEN '[0-5>'
                    WHEN `duration` >= 300 AND `duration` < 600 THEN '[5-10>'
                    WHEN `duration` >= 600 AND `duration` < 900 THEN '[10-15>'
                    ELSE '15 or more'
                    END `bin`
         FROM `sessions`
     ) `b`
     ON `a`.`bin` = `b`.`bin`
GROUP BY `a`.`bin`
```

# 1440.evaluate-boolean-expression 连表比较 
```sql
# Link: https://leetcode-cn.com/problems/evaluate-boolean-expression

SELECT `e`.`left_operand`,
       `operator`,
       `e`.`right_operand`,
       CASE
           WHEN `v1`.`value` > `v2`.`value` AND `operator` = '>' THEN 'true'
           WHEN `v1`.`value` < `v2`.`value` AND `operator` = '<' THEN 'true'
           WHEN `v1`.`value` = `v2`.`value` AND `operator` = '=' THEN 'true'
           ELSE 'false'
           END `value`
FROM `expressions`        `e`
         JOIN `variables` `v1`
         JOIN `variables` `v2`
              ON `e`.`left_operand` = `v1`.`name` AND `e`.`right_operand` = `v2`.`name`
```

# 1445.apples-oranges 销量差 
```sql
# Link: https://leetcode-cn.com/problems/apples-oranges

SELECT `sale_date`,
       SUM(IF(`fruit` = `apples`, `sold_num`, -`sold_num`)) `diff`
FROM `sales`
GROUP BY `sale_date`
ORDER BY `sale_date`
```

# 1454.active-users 连续7天在线的用户 
```sql
# Link: https://leetcode-cn.com/problems/active-users


SELECT `t3`.`id`, `name`
FROM (
         SELECT DISTINCT `id`
         FROM (
                  SELECT `id`, SUBDATE(`login_date`, `r`) `diff`
                  FROM (
                           SELECT `id`,
                                  `login_date`,
                                  ROW_NUMBER() OVER (PARTITION BY `id` ORDER BY `login_date`) `r`
                           FROM (
                                    SELECT DISTINCT *
                                    FROM `logins`) `l`) `t1`
                  GROUP BY `id`, `diff`
                  HAVING COUNT(*) >= 5) `t2`) `t3`
         JOIN `accounts` USING (`id`)
ORDER BY `id`

```

# 1459.rectangles-area 组合成矩形 
```sql
# Link: https://leetcode-cn.com/problems/rectangles-area


SELECT `a`.`id`                                                               `p1`,
       `b`.`id`                                                               `p2`,
       ABS((`a`.`x_value` - `b`.`x_value`) * (`a`.`y_value` - `b`.`y_value`)) `area`
FROM `points`          `a`
         JOIN `points` `b`
WHERE `a`.`id` < `b`.`id`
HAVING `area` > 0
ORDER BY `area` DESC, `p1`, `p2`;

```

# 1468.calculate-salaries 扣税 
```sql
# Link: https://leetcode-cn.com/problems/calculate-salaries


SELECT `s`.`company_id`,
       `s`.`employee_id`,
       `s`.`employee_name`,
       (
           CASE
               WHEN `maxsalary` < 1000 THEN `salary`
               WHEN `maxsalary` >= 1000 AND `maxsalary` < 10000 THEN ROUND(`salary` - `salary` * 0.24)
               WHEN `maxsalary` >= 10000 THEN ROUND(`salary` - `salary` * 0.49)
               END
           ) `salary`
FROM `salaries`      `s`
         LEFT JOIN (
                       SELECT `company_id`, MAX(`salary`) `maxsalary`
                       FROM `salaries`
                       GROUP BY `company_id`
                   ) `m`
                   ON `m`.`company_id` = `s`.`company_id`

```

# 1479.sales-by-day-of-the-week 每周销量统计 
```sql
# Link: https://leetcode-cn.com/problems/sales-by-day-of-the-week


SELECT DISTINCT `i`.`item_category`                                         `category`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 2, `o`.`quantity`, 0)) `monday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 3, `o`.`quantity`, 0)) `tuesday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 4, `o`.`quantity`, 0)) `wednesday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 5, `o`.`quantity`, 0)) `thursday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 6, `o`.`quantity`, 0)) `friday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 7, `o`.`quantity`, 0)) `saturday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 1, `o`.`quantity`, 0)) `sunday`
FROM `items`                `i`
         LEFT JOIN `orders` `o` ON `i`.`item_id` = `o`.`item_id`
GROUP BY `category`
ORDER BY `category`

```

# 1485.group-sold-products-by-the-date group聚合 
```sql
# Link: https://leetcode-cn.com/problems/group-sold-products-by-the-date

SELECT `sell_date`,
       COUNT(DISTINCT `product`)        `num_sold`,
       GROUP_CONCAT(DISTINCT `product`) `products`
FROM `activities`
GROUP BY `sell_date`
ORDER BY `sell_date`
```

# 1495.friendly-movies-streamed-last-month 儿童适宜的电影 
```sql
# Link: https://leetcode-cn.com/problems/friendly-movies-streamed-last-month

SELECT DISTINCT `title`
FROM `tvprogram`             `t`
         LEFT JOIN `content` `c` ON `t`.`content_id` = `c`.`content_id`
WHERE LEFT(`t`.`program_date`, 7) = '2020-06'
  AND `c`.`kids_content` = 'Y'
  AND `c`.`content_type` = 'Movies';

```

# 1501.countries-you-can-safely-invest-in 通话时长高于平均 
```sql
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

```

# 1511.customer-order-frequency 六七月消费大于等于100 
```sql
# Link: https://leetcode-cn.com/problems/customer-order-frequency

SELECT `c`.`customer_id`, `c`.`name`
FROM `customers`        `c`
         JOIN `orders`  `o` ON `o`.`customer_id` = `c`.`customer_id`
         JOIN `product` `p` ON `p`.`product_id` = `o`.`product_id`
GROUP BY `c`.`customer_id`, `c`.`name`
HAVING SUM(IF(LEFT(`o`.`order_date`, 7) = '2020-06', `p`.`price` * `o`.`quantity`, 0)) >= 100
   AND SUM(IF(LEFT(`o`.`order_date`, 7) = '2020-07', `p`.`price` * `o`.`quantity`, 0)) >= 100;
```

# 1517.find-users-with-valid-e-mails 邮箱正则 
```sql
# Link: https://leetcode-cn.com/problems/find-users-with-valid-e-mails

SELECT *
FROM `users`
WHERE `mail` REGEXP '^[a-zA-Z][a-zA-Z0-9\_\.\-]*@leetcode\.com$'
```

# 1527.patients-with-a-condition 患病患者 
```sql
# Link: https://leetcode-cn.com/problems/patients-with-a-condition

# Write your MySQL query statement below

SELECT *
FROM `patients`
WHERE `conditions` LIKE 'DIAB1%'
   OR `conditions` LIKE '% DIAB1%'
```

# 1532.the-most-recent-three-orders 最近三笔订单 
```sql
# Link: https://leetcode-cn.com/problems/the-most-recent-three-orders


SELECT `name` `customer_name`,
       `customer_id`,
       `order_id`,
       `order_date`
FROM (
         SELECT `c`.`name`,
                `c`.`customer_id`,
                `o`.`order_id`,
                `o`.`order_date`,
                ROW_NUMBER() OVER (PARTITION BY `c`.`customer_id` ORDER BY `o`.`order_date` DESC) `r`
         FROM `customers`       `c`
                  JOIN `orders` `o` ON `c`.`customer_id` = `o`.`customer_id`
     ) AS `a`
WHERE `r` <= 3
ORDER BY `customer_name`, `customer_id`, `order_date` DESC
```

# 1543.fix-product-name-format 格式化 
```sql
# Link: https://leetcode-cn.com/problems/fix-product-name-format


SELECT LOWER(TRIM(`product_name`)) `product_name`,
       LEFT(`sale_date`, 7)        `sale_date`,
       COUNT(*)                    `total`
FROM `sales`
GROUP BY LOWER(TRIM(`product_name`)), LEFT(`sale_date`, 7)
ORDER BY LOWER(TRIM(`product_name`)), LEFT(`sale_date`, 7)

```

# 1549.the-most-recent-orders-for-each-product 产品最新订单 
```sql
# Link: https://leetcode-cn.com/problems/the-most-recent-orders-for-each-product

SELECT `product_name`, `p`.`product_id`, `order_id`, `order_date`
FROM `products`          `p`
         JOIN
     (SELECT `o`.`product_id`,
             `o`.`order_id`,
             `o`.`order_date`,
             RANK() OVER (PARTITION BY `o`.`product_id` ORDER BY `o`.`order_date` DESC) `r`
      FROM `orders` `o`) `t` ON `p`.`product_id` = `t`.`product_id`
WHERE `r` = 1
ORDER BY `p`.`product_name`, `p`.`product_id`, `t`.`order_id`;

```

# 1555.bank-account-summary 是否透支 
```sql
# Link: https://leetcode-cn.com/problems/bank-account-summary


SELECT `u`.`user_id`,
       `u`.`user_name`,
       `credit` + IFNULL(`t`.`amount`, 0)                       `credit`,
       IF(`credit` + IFNULL(`t`.`amount`, 0) >= 0, 'No', 'Yes') `credit_limit_breached`
FROM `users`                            `u`
         LEFT JOIN (SELECT `user_id`, SUM(`amount`) `amount`
                    FROM (
                             (
                                 SELECT `paid_by` `user_id`, -SUM(`amount`) `amount`
                                 FROM `transactions`
                                 GROUP BY `paid_by`
                                 UNION
                                 SELECT `paid_to` `user_id`, SUM(`amount`) `amount`
                                 FROM `transactions`
                                 GROUP BY `paid_to`)
                         ) `t1`
                    GROUP BY `user_id`) `t`
                   ON `u`.`user_id` = `t`.`user_id`
GROUP BY `user_id`
```

# 1565.unique-orders-and-customers-per-month 月订单数和顾客数 
```sql
# Link: https://leetcode-cn.com/problems/unique-orders-and-customers-per-month

SELECT LEFT(`o`.`order_date`, 7)     `month`,
       COUNT(`order_id`)             `order_count`,
       COUNT(DISTINCT `customer_id`) `customer_count`
FROM `orders` `o`
WHERE `o`.`invoice` > 20
GROUP BY LEFT(`o`.`order_date`, 7);
```

# 1571.warehouse-manager 计算立方 
```sql
# Link: https://leetcode-cn.com/problems/warehouse-manager


SELECT `w`.`name`                                                   `warehouse_name`,
       SUM(`w`.`units` * `p`.`width` * `p`.`length` * `p`.`height`) `volume`
FROM `warehouse` `w`
         JOIN
     `products`  `p`
     ON `w`.`product_id` = `p`.`product_id`
GROUP BY `w`.`name`;

```

# 1581.customer-who-visited-but-did-not-make-any-transactions 进店未交易 
```sql
# Link: https://leetcode-cn.com/problems/customer-who-visited-but-did-not-make-any-transactions

SELECT `v`.`customer_id`,
       COUNT(DISTINCT `v`.`visit_id`) `count_no_trans`
FROM `visits`                     `v`
         LEFT JOIN `transactions` `t` ON `v`.`visit_id` = `t`.`visit_id`
WHERE `t`.`visit_id` IS NULL
GROUP BY `v`.`customer_id`


```

# 1587.bank-account-summary-ii 余大于10000的用户 
```sql
# Link: https://leetcode-cn.com/problems/bank-account-summary-ii


SELECT `name`,
       SUM(`amount`) `balance`
FROM `users`,
     `transactions`
WHERE `users`.`account` = `transactions`.`account`
GROUP BY `name`
HAVING SUM(`amount`) > 10000;



```

# 1596.the-most-frequently-ordered-products-for-each-customer 经常订购的商品 
```sql
# Link: https://leetcode-cn.com/problems/the-most-frequently-ordered-products-for-each-customer

SELECT `o`.`customer_id`, `o`.`product_id`, `p`.`product_name`
FROM (
         SELECT `customer_id`, `product_id`, RANK() OVER (PARTITION BY `customer_id` ORDER BY `c` DESC) `r`
         FROM (
                  SELECT `customer_id`, `product_id`, COUNT(*) `c`
                  FROM `orders`
                  GROUP BY `customer_id`, `product_id`
              ) `a`
     )                   `o`
         JOIN `products` `p`
              ON `o`.`product_id` = `p`.`product_id`
WHERE `o`.`r` = 1
ORDER BY `customer_id`, `product_id`
```

# 1607.sellers-with-no-sales 没有卖出产品的销售 
```sql
# Link: https://leetcode-cn.com/problems/sellers-with-no-sales


SELECT `seller_name`
FROM `seller`
WHERE `seller_id` NOT IN (SELECT DISTINCT `seller_id`
                          FROM `orders`
                          WHERE `sale_date` BETWEEN '2020-01-01' AND '2020-12-31')
ORDER BY `seller_name`;

SELECT `seller_name`
FROM `seller` `s`
WHERE NOT EXISTS(SELECT DISTINCT `seller_id`
                 FROM `orders` `o`
                 WHERE `sale_date` BETWEEN '2020-01-01' AND '2020-12-31'
                   AND `o`.`seller_id` = `s`.`seller_id`)
ORDER BY `seller_name`;
```

# 1613.find-the-missing-ids 缺失的id 
```sql
# Link: https://leetcode-cn.com/problems/find-the-missing-ids

WITH RECURSIVE `num`(`n`) AS (
    SELECT 1 `a`
    UNION ALL
    SELECT `n` + 1
    FROM `num`
    WHERE `n` < 100
)
SELECT `n` `ids`
FROM `num`
WHERE `n` NOT IN (
    SELECT `customer_id`
    FROM `customers`
)
  AND `n` <= (
    SELECT MAX(`customer_id`)
    FROM `customers`
)
```

# 1623.all-valid-triplets-that-can-represent-a-country 三个学校组合不同的代表队 
```sql
# Link: https://leetcode-cn.com/problems/all-valid-triplets-that-can-represent-a-country

SELECT `a`.`student_name` `member_a`,
       `b`.`student_name` `member_b`,
       `c`.`student_name` `member_c`
FROM `schoola` `a`,
     `schoolb` `b`,
     `schoolc` `c`
WHERE `a`.`student_id` != `b`.`student_id`
  AND `b`.`student_id` != `c`.`student_id`
  AND `a`.`student_id` != `c`.`student_id`
  AND `a`.`student_name` != `b`.`student_name`
  AND `b`.`student_name` != `c`.`student_name`
  AND `a`.`student_name` != `c`.`student_name`
```

# 1633.percentage-of-users-attended-a-contest 注册率 
```sql
# Link: https://leetcode-cn.com/problems/percentage-of-users-attended-a-contest


SELECT `contest_id`,
       ROUND(COUNT(DISTINCT `r`.`user_id`) / COUNT(DISTINCT `u`.`user_id`) * 100, 2) `percentage`
FROM `register`       `r`
         JOIN `users` `u`
GROUP BY `contest_id`
ORDER BY `percentage` DESC, `contest_id`
```

# 1635.hopper-company-queries-i 活跃数与行程数 
```sql
# Link: https://leetcode-cn.com/problems/hopper-company-queries-i

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `t1`.`month`,
       SUM(IFNULL(`c`, 0)) OVER (ORDER BY `t1`.`month`) `active_drivers`,
       IFNULL(`active_rides`, 0)                        `accepted_rides`
FROM `t1`
         LEFT JOIN
     (
         SELECT IF(YEAR(`join_date`) < '2020', 1, MONTH(`join_date`)) `month`, COUNT(`driver_id`) `c`
         FROM `drivers`
         WHERE YEAR(`join_date`) <= '2020'
         GROUP BY `month`
     ) `t2` ON `t1`.`month` = `t2`.`month`
         LEFT JOIN
     (
         SELECT MONTH(`r`.`requested_at`) `month`, COUNT(`driver_id`) `active_rides`
         FROM `acceptedrides`       `a`
                  LEFT JOIN `rides` `r` ON `a`.`ride_id` = `r`.`ride_id`
             AND YEAR(`r`.`requested_at`) >= '2020' AND YEAR(`r`.`requested_at`) < '2021'
         GROUP BY `month`
     ) `t3` ON `t1`.`month` = `t3`.`month`
```

# 1645.hopper-company-queries-ii 工作率 
```sql
# Link: https://leetcode-cn.com/problems/hopper-company-queries-ii

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `t1`.`month`,
       IFNULL(ROUND(IFNULL(`active_rides`, 0) / SUM(IFNULL(`c`, 0)) OVER (ORDER BY `t1`.`month`) * 100, 2),
              0) `working_percentage`
FROM `t1`
         LEFT JOIN
     (
         SELECT IF(YEAR(`join_date`) < '2020', 1, MONTH(`join_date`)) `month`, COUNT(DISTINCT `driver_id`) `c`
         FROM `drivers`
         WHERE YEAR(`join_date`) <= '2020'
         GROUP BY `month`
     ) `t2`
     ON `t1`.`month` = `t2`.`month`
         LEFT JOIN
     (
         SELECT MONTH(`r`.`requested_at`) `month`, COUNT(DISTINCT `driver_id`) `active_rides`
         FROM `acceptedrides`       `a`
                  LEFT JOIN `rides` `r` ON `a`.`ride_id` = `r`.`ride_id`
             AND YEAR(`r`.`requested_at`) >= '2020' AND YEAR(`r`.`requested_at`) < '2021'
         GROUP BY `month`
     ) `t3` ON `t1`.`month` = `t3`.`month`
```

# 1651.hopper-company-queries-iii 平均距离和平均时间 
```sql
# Link: https://leetcode-cn.com/problems/hopper-company-queries-iii

WITH RECURSIVE `t1`(`month`) AS (
    SELECT 1
    UNION ALL
    SELECT `month` + 1
    FROM `t1`
    WHERE `month` + 1 <= 12
)

SELECT `month`, `average_ride_distance`, `average_ride_duration`
FROM (SELECT `month`,
             ROUND(AVG(IFNULL(`ride_distance`, 0)) OVER (ORDER BY `month` ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING ),
                   2)                             `average_ride_distance`,
             ROUND(AVG(IFNULL(`ride_duration`, 0)) OVER (ORDER BY `month` ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING ),
                   2)                             `average_ride_duration`,
             ROW_NUMBER() OVER (ORDER BY `month`) `r`
      FROM `t1`
               LEFT JOIN
           (SELECT DISTINCT MONTH(`requested_at`)                                                        `month`,
                            SUM(`ride_distance`)
                                OVER (PARTITION BY MONTH(`requested_at`) ORDER BY MONTH(`requested_at`)) `ride_distance`,
                            SUM(`ride_duration`)
                                OVER (PARTITION BY MONTH(`requested_at`) ORDER BY MONTH(`requested_at`)) `ride_duration`
            FROM `acceptedrides`
                     JOIN `rides` USING (`ride_id`)
            WHERE YEAR(`requested_at`) = 2020) `a`
           USING (`month`)) `b`
WHERE `b`.`r` <= 10

```

# 1661.average-time-of-process-per-machine 平均运行时间 
```sql
# Link: https://leetcode-cn.com/problems/average-time-of-process-per-machine

SELECT `machine_id`,
       ROUND(SUM(IF(`activity_type` = 'end', `timestamp`, -`timestamp`)) / COUNT(*) * 2, 3) `processing_time`
FROM `activity`
GROUP BY `machine_id`
```

# 1667.fix-names-in-a-table  修复名字 
```sql
# Link: https://leetcode-cn.com/problems/fix-names-in-a-table

SELECT `user_id`,
       CONCAT(UPPER(LEFT(`name`, 1)), LOWER(SUBSTR(`name`, 2))) `name`
FROM `users`
ORDER BY `user_id`
```

# 1677.products-worth-over-invoices 统计 
```sql
# Link: https://leetcode-cn.com/problems/products-worth-over-invoices

SELECT `p`.`name`                     `name`,
       IFNULL(SUM(`i`.`rest`), 0)     `rest`,
       IFNULL(SUM(`i`.`paid`), 0)     `paid`,
       IFNULL(SUM(`i`.`canceled`), 0) `canceled`,
       IFNULL(SUM(`i`.`refunded`), 0) `refunded`
FROM `product`               `p`
         LEFT JOIN `invoice` `i`
                   ON `p`.`product_id` = `i`.`product_id`
GROUP BY `p`.`product_id`
ORDER BY `p`.`name`
```

# 1683.invalid-tweets 无效推文 
```sql
# Link: https://leetcode-cn.com/problems/invalid-tweets

SELECT `tweet_id`
FROM `tweets`
WHERE CHAR_LENGTH(`content`) > 15
```

# 1693.daily-leads-and-partners 统计 
```sql
# Link: https://leetcode-cn.com/problems/daily-leads-and-partners

SELECT `date_id`,
       `make_name`,
       COUNT(DISTINCT `lead_id`)    `unique_leads`,
       COUNT(DISTINCT `partner_id`) `unique_partners`
FROM `dailysales`
GROUP BY `date_id`, `make_name`

```

# 1699.number-of-calls-between-two-persons 通话次数 
```sql
# Link: https://leetcode-cn.com/problems/number-of-calls-between-two-persons

SELECT `person1`, `person2`, COUNT(1) `call_count`, SUM(`duration`) `total_duration`
FROM (
         SELECT IF(`from_id` > `to_id`, `to_id`, `from_id`) `person1`,
                IF(`from_id` > `to_id`, `from_id`, `to_id`) `person2`,
                `duration`
         FROM `calls`) `c`
GROUP BY `person1`, `person2`;

SELECT `from_id`        `person1`,
       `to_id`          `person2`,
       COUNT(`from_id`) `call_count`,
       SUM(`duration`)  `total_duration`
FROM `calls`
GROUP BY LEAST(`from_id`, `to_id`), GREATEST(`from_id`, `to_id`)
```

# 1709.biggest-window-between-visits 日期缺失 
```sql
# Link: https://leetcode-cn.com/problems/biggest-window-between-visits


SELECT `user_id`, MAX(`diff`) `biggest_window`
FROM (
         SELECT `user_id`,
                DATEDIFF(
                        LEAD(`visit_date`, 1, '2021-01-01') OVER (PARTITION BY `user_id` ORDER BY `visit_date`),
                        `visit_date`
                    ) `diff`
         FROM `uservisits`
     ) `tmp`
GROUP BY `user_id`


```

# 1715.count-apples-and-oranges 统计 
```sql
# Link: https://leetcode-cn.com/problems/count-apples-and-oranges


SELECT SUM(`b`.`apple_count` + IFNULL(`c`.`apple_count`, 0))   `apple_count`,
       SUM(`b`.`orange_count` + IFNULL(`c`.`orange_count`, 0)) `orange_count`
FROM `boxes`                `b`
         LEFT JOIN `chests` `c` ON `b`.`chest_id` = `c`.`chest_id`
```

# 1729.find-followers-count 关注数 
```sql
# Link: https://leetcode-cn.com/problems/find-followers-count

SELECT `user_id`, COUNT(`follower_id`) `followers_count`
FROM `followers`
GROUP BY `user_id`
ORDER BY `user_id`


```

# 1731.the-number-of-employees-which-report-to-each-employee 每位经理有多少下属 
```sql
# Link: https://leetcode-cn.com/problems/the-number-of-employees-which-report-to-each-employee

SELECT `e1`.`reports_to`        `employee_id`,
       `e2`.`name`,
       COUNT(`e1`.`reports_to`) `reports_count`,
       ROUND(AVG(`e1`.`age`))   `average_age`
FROM `employees`          `e1`
         JOIN `employees` `e2` ON `e1`.`reports_to` = `e2`.`employee_id`
GROUP BY `e1`.`reports_to`, `e2`.`name`
ORDER BY `e1`.`reports_to`
```

# 1741.find-total-time-spent-by-each-employee 员工花费时间 
```sql
# Link: https://leetcode-cn.com/problems/find-total-time-spent-by-each-employee

SELECT `event_day` `day`, `emp_id`, SUM(`out_time` - `in_time`) `total_time`
FROM `employees`
GROUP BY `day`, `emp_id`
```

# 1747.leetflex-banned-accounts 应被禁止的账户 
```sql
# Link: https://leetcode-cn.com/problems/leetflex-banned-accounts

SELECT DISTINCT `l1`.`account_id`
FROM `loginfo`          `l1`
         JOIN `loginfo` `l2`
              ON `l1`.`account_id` = `l2`.`account_id`
                  AND `l1`.`ip_address` <> `l2`.`ip_address`
                  AND ((`l1`.`login` BETWEEN `l2`.`login` AND `l2`.`logout`)
                      OR (`l1`.`logout` BETWEEN `l2`.`login` AND `l2`.`logout`))

```

# 1757.recyclable-and-low-fat-products 可回收且低脂的产品 
```sql
# Link: https://leetcode-cn.com/problems/recyclable-and-low-fat-products

SELECT `product_id`
FROM `products`
WHERE `low_fats` = 'Y'
  AND `recyclable` = 'Y'
```

# 1767.find-the-subtasks-that-did-not-execute 找到未执行的子任务 
```sql
# Link: https://leetcode-cn.com/problems/find-the-subtasks-that-did-not-execute

WITH RECURSIVE `t`(`task_id`, `subtask_id`) AS (
    SELECT `task_id`, `subtasks_count`
    FROM `tasks`
    UNION ALL
    SELECT `task_id`, `subtask_id` - 1
    FROM `t`
    WHERE `subtask_id` - 1 > 0
)

SELECT *
FROM `t`
         LEFT JOIN `executed` USING (`task_id`, `subtask_id`)
WHERE `executed`.`subtask_id` IS NULL
ORDER BY `task_id`, `subtask_id`
```

# 1777.products-price-for-each-store 每家商店价格 
```sql
# Link: https://leetcode-cn.com/problems/products-price-for-each-store

SELECT `product_id`,
       SUM(IF(`store` = 'store1', `price`, NULL)) `store1`,
       SUM(IF(`store` = 'store2', `price`, NULL)) `store2`,
       SUM(IF(`store` = 'store3', `price`, NULL)) `store3`
FROM `products`
GROUP BY `product_id`;
```

# 1783.grand-slam-titles 大满贯数 
```sql
# Link: https://leetcode-cn.com/problems/grand-slam-titles


WITH `t` AS (SELECT `wimbledon` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `fr_open` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `us_open` `tournament`
             FROM `championships`
             UNION ALL
             SELECT `au_open` `tournament`
             FROM `championships`)

SELECT `p`.`player_id`, `p`.`player_name`, COUNT(*) `grand_slams_count`
FROM `t`
         LEFT JOIN `players` `p` ON `t`.`tournament` = `p`.`player_id`
GROUP BY `p`.`player_id`, `p`.`player_name`
```

# 1789.primary-department-for-each-employee 员工的主要部门 
```sql
# Link: https://leetcode-cn.com/problems/primary-department-for-each-employee

SELECT DISTINCT `employee_id`,
                FIRST_VALUE(`department_id`) OVER (PARTITION BY `employee_id` ORDER BY `primary_flag`) `department_id`
FROM `employee`
```

# 1795.rearrange-products-table 商品价格合并 
```sql
# Link: https://leetcode-cn.com/problems/rearrange-products-table

SELECT `product_id`, 'store1' `store`, `store1` `price`
FROM `products`
WHERE `store1` IS NOT NULL
UNION ALL
SELECT `product_id`, 'store2' `store`, `store2` `price`
FROM `products`
WHERE `store2` IS NOT NULL
UNION ALL
SELECT `product_id`, 'store3' `store`, `store3` `price`
FROM `products`
WHERE `store3` IS NOT NULL

```

# 1809.ad-free-sessions 没有展示广告的会话 
```sql
# Link: https://leetcode-cn.com/problems/ad-free-sessions

SELECT `session_id`
FROM `playback`
WHERE `session_id` NOT IN (
    SELECT DISTINCT `session_id`
    FROM `playback`     `p`
             JOIN `ads` `a` ON `p`.`customer_id` = `a`.`customer_id`
    WHERE `timestamp` BETWEEN `start_time` AND `end_time`)
```

# 1811.find-interview-candidates 统计 
```sql
# Link: https://leetcode-cn.com/problems/find-interview-candidates

WITH `a` AS
         (SELECT DISTINCT `temp1`.`name`, `mail`
          FROM (
                   SELECT `contest_id`,
                          `u1`.`user_id`,
                          `u1`.`name`,
                          `mail`,
                          `contest_id` - ROW_NUMBER() OVER (PARTITION BY `u1`.`user_id` ORDER BY `contest_id`) `diffs`
                   FROM `users`             `u1`
                            JOIN `contests` `c`
                                 ON `u1`.`user_id` IN (`c`.`gold_medal`, `c`.`silver_medal`, `c`.`bronze_medal`)
                   ORDER BY `u1`.`user_id`, `contest_id`) `temp1`
          GROUP BY `temp1`.`user_id`, `diffs`
          HAVING COUNT(`contest_id`) >= 3
          ORDER BY `temp1`.`user_id`)
SELECT *
FROM `a`
UNION
SELECT `u2`.`name`, `mail`
FROM `users`             `u2`
         JOIN `contests` `c1`
              ON `u2`.`user_id` = `c1`.`gold_medal`
GROUP BY `c1`.`gold_medal`
HAVING COUNT(*) >= 3


```

# 1821.find-customers-with-positive-revenue-this-year 正收入 
```sql
# Link: https://leetcode-cn.com/problems/find-customers-with-positive-revenue-this-year

SELECT `customer_id`
FROM `customers`
WHERE `year` = 2021
GROUP BY `customer_id`
HAVING SUM(`revenue`) > 0
```

# 1831.maximum-transaction-each-day 每天最大交易 
```sql
# Link: https://leetcode-cn.com/problems/maximum-transaction-each-day


SELECT `transaction_id`
FROM (
         SELECT `transaction_id`,
                RANK() OVER (PARTITION BY LEFT(`day`, 10) ORDER BY `amount` DESC) `r`
         FROM `transactions`
         ORDER BY `transaction_id`
     ) `t`
WHERE `r` = 1
```

