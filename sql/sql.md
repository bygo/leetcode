# 0175.combine-two-tables 
```sql
# Title: Combine Two Tables
# Link: https://leetcode-cn.com/problems/combine-two-tables

# 给所有人加上地址
SELECT `FirstName`, `LastName`, `City`, `State`
FROM `Person`
         LEFT JOIN `Address` ON `Person`.`PersonId` = `Address`.`PersonId`
```

# 0176.second-highest-salary 
```sql
# Title: Second Highest Salary
# Link: https://leetcode-cn.com/problems/second-highest-salary

# 找出第二名 Null 也返回
SELECT (
           SELECT DISTINCT `Salary`
           FROM `Employee`
           ORDER BY `Salary` DESC
           LIMIT 1,1) `SecondHighestSalary`
```

# 0177.nth-highest-salary 
```sql
# Title: Nth Highest Salary
# Link: https://leetcode-cn.com/problems/nth-highest-salary

# 找出第N名 Null 也返回
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N := N - 1;
    RETURN (
        SELECT (
                   SELECT DISTINCT `Salary`
                   FROM `Employee`
                   ORDER BY Salary DESC
                   LIMIT N,1)  `NthHighestSalary`
    );
END
```

# 0178.rank-scores 
```sql
# Title: Rank Scores
# Link: https://leetcode-cn.com/problems/rank-scores

SELECT `score`,
       @rank := @rank + (@pre != (@pre := `score`)) as `Rank`
from scores,
     (select @pre := -1, @rank := 0) as t
order by `score` desc
```

# 0180.consecutive-numbers 
```sql
# Title: Consecutive Numbers
# Link: https://leetcode-cn.com/problems/consecutive-numbers

SELECT DISTINCT `Num` `ConsecutiveNums`
FROM (SELECT IF(@pre = `Num`, @count := @count + 1, @count := 1) AS `counter`, @pre := `Num` AS `Num`
      FROM `Logs`,
           (SELECT @pre := 0) AS t) AS t
WHERE `counter` >= 3;
```

# 0180.consecutive-numbers 
```sql
# Title: Consecutive Numbers
# Link: https://leetcode-cn.com/problems/consecutive-numbers

SELECT DISTINCT `Num` `ConsecutiveNums`
FROM (SELECT IF(@pre = `Num`, @count := @count + 1, @count := 1) AS `counter`, @pre := `Num` AS `Num`
      FROM `Logs`,
           (SELECT @pre := 0) AS t) AS t
WHERE `counter` >= 3;
```

```sql

# Title: Consecutive Numbers
# Link: https://leetcode-cn.com/problems/consecutive-numbers

SELECT DISTINCT
    l1.Num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
```

# 0181.employees-earning-more-than-their-managers 
```sql
# Title: Employees Earning More Than Their Managers
# Link: https://leetcode-cn.com/problems/employees-earning-more-than-their-managers

SELECT E1.`Name` AS `Employee`
FROM `Employee` AS E1
         LEFT JOIN `Employee` AS E2 ON E1.ManagerId = E2.id
where E1.Salary > E2.Salary
```

# 0182.duplicate-emails 
```sql
# Title: Duplicate Emails
# Link: https://leetcode-cn.com/problems/duplicate-emails

SELECT `Email`
FROM `Person`
GROUP BY `Email`
HAVING count(`Email`) > 1x`x`
```

# 0183.customers-who-never-order 
```sql
# Title: Customers Who Never Order
# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT Customers.Name Customers
FROM Customers
         LEFT JOIN Orders  ON Customers.Id = Orders.CustomerId
WHERE Orders.CustomerId is null
```

# 0183.customers-who-never-order 
```sql
# Title: Customers Who Never Order
# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT Customers.Name Customers
FROM Customers
         LEFT JOIN Orders  ON Customers.Id = Orders.CustomerId
WHERE Orders.CustomerId is null
```

```sql
# Title: Customers Who Never Order
# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT `name` AS Customers
FROM `Customers`
WHERE `id` NOT IN (SELECT `CustomerId` FROM `Orders`)
```

# 0184.department-highest-salary 
```sql
# Title: Department Highest Salary
# Link: https://leetcode-cn.com/problems/department-highest-salary

SELECT Department.name AS 'Department',
       Employee.name   AS 'Employee',
       Salary
FROM Employee
         JOIN
     Department ON Employee.DepartmentId = Department.Id
WHERE (Employee.DepartmentId, Salary) IN
      (SELECT DepartmentId,
              MAX(Salary)
       FROM Employee
       GROUP BY DepartmentId
      )
```

# 0185.department-top-three-salaries 
```sql
# Title: Department Top Three Salaries
# Link: https://leetcode-cn.com/problems/department-top-three-salaries

SELECT
	Department.Name AS Department,
	e1.Name AS Employee,
	e1.Salary AS Salary
FROM
	Employee AS e1 JOIN Department
ON
	e1.DepartmentId = Department.Id
	AND  (SELECT count( DISTINCT e2.Salary )
			 FROM Employee AS e2
			 WHERE	e1.Salary < e2.Salary AND e1.DepartmentId = e2.DepartmentId )  < 3
ORDER BY Salary DESC;

```

# 0196.delete-duplicate-emails 
```sql
# Title: Delete Duplicate Emails
# Link: https://leetcode-cn.com/problems/delete-duplicate-emails

DELETE p1
FROM `Person` p1
         JOIN `Person` p2
WHERE p1.Email = p2.Email
  AND p1.Id > p2.Id
```

# 0197.rising-temperature 
```sql
# Title: Rising Temperature
# Link: https://leetcode-cn.com/problems/rising-temperature

SELECT w1.Id AS Id
FROM Weather w1
         JOIN Weather w2
WHERE DATEDIFF(w1.recordDate, w2.recordDate) = 1
  AND w1.Temperature > w2.Temperature
```

# 0262.trips-and-users 
```sql
# Title: Trips and Users
# Link: https://leetcode-cn.com/problems/trips-and-users

SELECT t.request_at AS `Day`,
       ROUND(
                   SUM(
                           IF(t.STATUS = 'completed', 0, 1)
                       )
                   /
                   COUNT(t.STATUS),
                   2
           )        AS `Cancellation Rate`
FROM Trips t,
     Users u1,
     Users u2
WHERE t.client_id = u1.users_id
  AND u1.banned = 'No'
  AND t.driver_id = u2.users_id
  AND u2.banned = 'No'
  AND t.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY t.request_at
```

# 0511.game-play-analysis-i 
```sql
# Title: Game Play Analysis I
# Link: https://leetcode-cn.com/problems/game-play-analysis-i

SELECT `player_id`,  `device_id`
FROM Activity order by `event_date`
GROUP BY `player_id`

```

# 0512.game-play-analysis-ii 
```sql
# Title: Game Play Analysis II
# Link: https://leetcode-cn.com/problems/game-play-analysis-ii

SELECT `player_id`, `device_id`
FROM Activity
WHERE (`player_id`, `event_date`) IN (SELECT `player_id`, MIN(`event_date`)
                                      FROM Activity
                                      GROUP BY `player_id`)
```

# 0534.game-play-analysis-iii 
```sql
# Title: Game Play Analysis III
# Link: https://leetcode-cn.com/problems/game-play-analysis-iii

SELECT a2.`player_id`         AS `player_id`,
       a2.`event_date`        AS `event_date`,
       SUM(a1.`games_played`) AS `games_played_so_far`
FROM Activity a1,
     Activity a2
WHERE a1.`player_id` = a2.`player_id`
  AND a1.event_date <= a2.event_date
GROUP BY a2.player_id, a2.`event_date`
```

# 0550.game-play-analysis-iv 
```sql
# Title: Game Play Analysis IV
# Link: https://leetcode-cn.com/problems/game-play-analysis-iv

SELECT ROUND(COUNT(a2.player_id) / COUNT(a1.player_id), 2) AS `fraction`
FROM (SELECT player_id, MIN(event_date) AS event_date
      FROM Activity
      GROUP BY player_id) a1
         LEFT JOIN Activity a2
                   ON a1.player_id = a2.player_id
                       AND DATEDIFF(a2.event_date, a1.event_date) = 1;
```

# 0569.median-employee-salary 
```sql
# Title: Median Employee Salary
# Link: https://leetcode-cn.com/problems/median-employee-salary

SELECT `id`, `Company`, `Salary`
FROM (SELECT `id`,
             `Company`,
             `Salary`,
             ROW_NUMBER() OVER (PARTITION BY `Company` ORDER BY `Salary`) AS `rank`,
             COUNT(*) OVER (PARTITION BY `Company`)                       AS `count`
      FROM `Employee`) t
WHERE `RANK` BETWEEN `count` / 2 AND `count` / 2 + 1
```

# 0570.managers-with-at-least-5-direct-reports 
```sql
# Title: Managers with at Least 5 Direct Reports
# Link: https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports

SELECT `Name`
FROM `employee` t1
         JOIN (SELECT `ManagerId` FROM `employee` GROUP BY `ManagerId` HAVING COUNT(*) >= 5) t2 ON t1.Id = t2.ManagerId
```

# 0571.find-median-given-frequency-of-numbers 
```sql
# Title: Find Median Given Frequency of Numbers
# Link: https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers


SELECT AVG(`Number`) AS Median
FROM (SELECT *,
             SUM(`Frequency`) OVER (ORDER BY `Number` ASC)  AS n1,
             SUM(`Frequency`) OVER (ORDER BY `Number` DESC) AS n2
      FROM `Numbers`) AS t
WHERE n1 BETWEEN n2 - `Frequency` AND n2 + `Frequency`;

```

# 0574.winning-candidate 
```sql
# Title: Winning Candidate
# Link: https://leetcode-cn.com/problems/winning-candidate

SELECT `Name`
FROM (SELECT `CandidateId` FROM vote GROUP BY CandidateId ORDER BY COUNT(`CandidateId`) DESC LIMIT 1) t
         JOIN `Candidate` ON Candidate.id = CandidateId
```

# 0577.employee-bonus 
```sql
# Title: Employee Bonus
# Link: https://leetcode-cn.com/problems/employee-bonus

SELECT `name`, `bonus`
FROM Employee
         LEFT JOIN Bonus ON Employee.empid = Bonus.empid
WHERE bonus.bonus < 1000
   OR bonus IS NULL
```

# 0578.get-highest-answer-rate-question 
```sql
# Title: Get Highest Answer Rate Question
# Link: https://leetcode-cn.com/problems/get-highest-answer-rate-question

SELECT `question_id` AS survey_log
FROM `survey_log`
GROUP BY `question_id`
ORDER BY SUM(IF(`action` = 'answer', 1, 0)) / SUM(IF(`action` = 'show', 1, 0)) DESC
LIMIT 1
```

# 0579.find-cumulative-salary-of-an-employee 
```sql
# Title: Find Cumulative Salary of an Employee
# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT E1.id,
       E1.month,
       (IFNULL(E1.salary, 0) + IFNULL(E2.salary, 0) + IFNULL(E3.salary, 0)) AS Salary
FROM (SELECT id,
             MAX(month) AS month
      FROM Employee
      GROUP BY id
      HAVING COUNT(*) > 1) AS maxmonth
         LEFT JOIN
     Employee E1 ON maxmonth.id = E1.id
         AND maxmonth.month > E1.month
         LEFT JOIN
     Employee E2 ON E2.id = E1.id
         AND E2.month = E1.month - 1
         LEFT JOIN
     Employee E3 ON E3.id = E1.id
         AND E3.month = E1.month - 2
ORDER BY id, month DESC
```

# 0579.find-cumulative-salary-of-an-employee 
```sql
# Title: Find Cumulative Salary of an Employee
# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT E1.id,
       E1.month,
       (IFNULL(E1.salary, 0) + IFNULL(E2.salary, 0) + IFNULL(E3.salary, 0)) AS Salary
FROM (SELECT id,
             MAX(month) AS month
      FROM Employee
      GROUP BY id
      HAVING COUNT(*) > 1) AS maxmonth
         LEFT JOIN
     Employee E1 ON maxmonth.id = E1.id
         AND maxmonth.month > E1.month
         LEFT JOIN
     Employee E2 ON E2.id = E1.id
         AND E2.month = E1.month - 1
         LEFT JOIN
     Employee E3 ON E3.id = E1.id
         AND E3.month = E1.month - 2
ORDER BY id, month DESC
```

```sql
# Title: Find Cumulative Salary of an Employee
# Link: https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee


SELECT id, Month, SUM(Salary) AS Salary
FROM (
         SELECT e1.id, e1.month, e1.salary
         FROM (SELECT `id`, `month`, `salary`, MAX(month) OVER (PARTITION BY `id`,`month`) AS maxmonth
               FROM `Employee`) e1
                  JOIN Employee e2
         WHERE e1.month != e1.maxmonth
           AND e1.id = e2.id
           AND e2.month BETWEEN e1.month - 2 AND e1.month
         ORDER BY `id`
     ) t
GROUP BY id, month
ORDER BY id, month DESC;

```

# 0580.count-student-number-in-departments 
```sql
# Title: Count Student Number in Departments
# Link: https://leetcode-cn.com/problems/count-student-number-in-departments

SELECT dept_name,COUNT(student_id) as student_number
FROM department d
         LEFT JOIN student s ON d.dept_id = s.dept_id
GROUP BY dept_name
ORDER BY student_number desc
```

# 0584.find-customer-referee 
```sql
# Title: Find Customer Referee
# Link: https://leetcode-cn.com/problems/find-customer-referee


SELECT `name`
FROM `customer`
WHERE referee_id != 2
   OR referee_id IS NULL
```

# 0585.investments-in-2016 
```sql
# Title: Investments in 2016
# Link: https://leetcode-cn.com/problems/investments-in-2016

SELECT ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM (SELECT *,
             COUNT(*) OVER ( PARTITION BY tiv_2015) AS y,
             COUNT(*) OVER ( PARTITION BY lat,lon)  AS p
      FROM insurance) t
WHERE y > 1 && p = 1
```

# 0586.customer-placing-the-largest-number-of-orders 
```sql
# Title: Customer Placing the Largest Number of Orders
# Link: https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders

SELECT customer_number
FROM orders
GROUP BY customer_number
ORDER BY COUNT(*) desc
LIMIT 1
```

# 0595.big-countries 
```sql
# Title: Big Countries
# Link: https://leetcode-cn.com/problems/big-countries

SELECT `name`, `population`, `area`
FROM `world`
WHERE area > 3000000
   OR population > 25000000
```

# 0596.classes-more-than-5-students 
```sql
# Title: Classes More Than 5 Students
# Link: https://leetcode-cn.com/problems/classes-more-than-5-students

SELECT class
FROM courses
GROUP BY class
HAVING COUNT(DISTINCT student) >= 5
```

# 0597.friend-requests-i-overall-acceptance-rate 
```sql
# Title: Friend Requests I: Overall Acceptance Rate
# Link: https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate

SELECT ROUND(
               IFNULL(
                           (SELECT COUNT(*) FROM (SELECT DISTINCT requester_id, accepter_id FROM RequestAccepted) AS A)
                           /
                           (SELECT COUNT(*) FROM (SELECT DISTINCT sender_id, send_to_id FROM FriendRequest) AS B),
                           0)
           , 2) AS accept_rate;
```

# 0601.human-traffic-of-stadium 
```sql
# Title: Human Traffic of Stadium
# Link: https://leetcode-cn.com/problems/human-traffic-of-stadium

WITH countT AS (SELECT id,
                       COUNT(*) OVER (PARTITION BY rn ORDER BY rn ) AS counter
                FROM (SELECT id,
                             id - ROW_NUMBER() OVER (ORDER BY id) AS rn
                      FROM stadium
                      WHERE people >= 100) rowT)
SELECT s.*
FROM stadium s
         JOIN countT ON
    s.id = countT.id
WHERE countT.counter > 2
ORDER BY s.visit_date;
```

# 0602.friend-requests-ii-who-has-the-most-friends 
```sql
# Title: Friend Requests II: Who Has the Most Friends
# Link: https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends

SELECT id, SUM(n) AS num
FROM (SELECT `accepter_id` AS id, COUNT(*) AS n
      FROM `request_accepted`
      GROUP BY accepter_id
      UNION ALL
      SELECT `requester_id` AS id, COUNT(*) AS n
      FROM `request_accepted`
      GROUP BY requester_id) t
GROUP BY id
ORDER BY num DESC
LIMIT 1;

```

# 0603.consecutive-available-seats 
```sql
# Title: Consecutive Available Seats
# Link: https://leetcode-cn.com/problems/consecutive-available-seats

SELECT `seat_id`
FROM (SELECT `seat_id`, COUNT(*) OVER (PARTITION BY r) AS c
      FROM (SELECT `seat_id`,
                   `seat_id` - ROW_NUMBER() OVER (ORDER BY `seat_id`) AS r
            FROM cinema
            WHERE free = 1) t) t2
WHERE c > 1
GROUP BY `seat_id`

```

# 0607.sales-person 
```sql
# Title: Sales Person
# Link: https://leetcode-cn.com/problems/sales-person

SELECT DISTINCT `name`
FROM salesperson
WHERE `sales_id` NOT IN (SELECT `sales_id`
                         FROM `orders`
                         WHERE com_id = (SELECT `com_id` FROM `company` WHERE `name` = 'RED'))
```

# 0608.tree-node 
```sql
# Title: Tree Node
# Link: https://leetcode-cn.com/problems/tree-node

SELECT DISTINCT t1.Id,
                IF(t1.p_id IS NULL, 'Root',
                   IF(t2.id IS NOT NULL, 'Inner', 'Leaf')) AS `Type`
FROM `tree` AS t1
         LEFT JOIN `tree` AS t2
                   ON t1.id = t2.p_id
```

# 0610.triangle-judgement 
```sql
# Title: Triangle Judgement
# Link: https://leetcode-cn.com/problems/triangle-judgement

SELECT *, IF(x + y > z AND x + z > y AND y + z > x, "Yes", "No") AS `triangle`
FROM `triangle`
```

# 0612.shortest-distance-in-a-plane 
```sql
# Title: Shortest Distance in a Plane
# Link: https://leetcode-cn.com/problems/shortest-distance-in-a-plane

SELECT ROUND(SQRT(MIN((POW(p1.x - p2.x, 2) + POW(p1.y - p2.y, 2)))), 2) AS shortest
FROM point_2d p1
         JOIN
     point_2d p2 ON (p1.x != p2.x AND p1.y = p2.y)
         OR (p1.x != p2.x AND p1.y != p2.y)
         OR (p1.x = p2.x AND p1.y != p2.y)

```

# 0613.shortest-distance-in-a-line 
```sql
# Title: Shortest Distance in a Line
# Link: https://leetcode-cn.com/problems/shortest-distance-in-a-line

SELECT MIN(ABS(p1.x - p2.x)) AS `shortest`
FROM point p1
         JOIN point p2 ON p1.x != p2.x;

```

# 0614.second-degree-follower 
```sql
# Title: Second Degree Follower
# Link: https://leetcode-cn.com/problems/second-degree-follower

select
    followee as follower, count(distinct follower) as num
from
    follow
where followee in (select follower from follow)
group by followee
order by followee



select
    followee as follower, count(distinct follower) as num
from
    follow
where followee in (select follower from follow)
group by followee
order by followee
```

# 0615.average-salary-departments-vs-company 
```sql
# Title: Average Salary: Departments VS Company
# Link: https://leetcode-cn.com/problems/average-salary-departments-vs-company

SELECT DISTINCT DATE_FORMAT(pay_date, '%Y-%m')                  AS pay_month,
                department_id,
                IF(b = a, 'same', IF(a < b, 'lower', 'higher')) AS comparison
FROM (SELECT department_id,
             amount,
             pay_date,
             AVG(amount) OVER (PARTITION BY pay_date,department_id) AS a,
             AVG(amount) OVER (PARTITION BY pay_date)               AS b
      FROM salary st
               JOIN employee et ON st.employee_id = et.employee_id) t
ORDER BY pay_month DESC;
```

# 0618.students-report-by-geography 
```sql
# Title: Students Report By Geography
# Link: https://leetcode-cn.com/problems/students-report-by-geography

SELECT America, Asia, Europe
FROM (SELECT `name`, ROW_NUMBER() OVER (ORDER BY name) AS r, name AS America
      FROM student
      WHERE continent = 'America') a
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY name) AS r, name AS Asia
                    FROM student
                    WHERE continent = 'Asia') b ON a.r = b.r
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY name) AS r, name AS Europe
                    FROM student
                    WHERE continent = 'Europe') c ON a.r = c.r
```

# 0619.biggest-single-number 
```sql
# Title: Biggest Single Number
# Link: https://leetcode-cn.com/problems/biggest-single-number

SELECT (SELECT num
        FROM my_numbers
        GROUP BY num
        HAVING COUNT(*) = 1
        ORDER BY num DESC
        LIMIT 1) num
```

# 0620.not-boring-movies 
```sql
# Title: Not Boring Movies
# Link: https://leetcode-cn.com/problems/not-boring-movies

SELECT `id`, `movie`, `description`, `rating`
FROM `cinema`
WHERE `description` != 'boring'
  AND id % 2 = 1
ORDER BY rating DESC
```

# 0626.exchange-seats 
```sql
# Title: Exchange Seats
# Link: https://leetcode-cn.com/problems/exchange-seats

select
    row_number() over(order by (id + 1 - 2 * power(0,id%2))) as id,
    student
from
    seat



SELECT
IF (id%2=0,id-1,id+1) as id
    student
FROM
ORDER BY id ASC;

```

# 0627.swap-salary 
```sql
# Title: Swap Salary
# Link: https://leetcode-cn.com/problems/swap-salary

UPDATE salary
SET `sex` = IF(`Sex` = 'f', 'm', 'f')
```

# 1045.customers-who-bought-all-products 
```sql
# Title: Customers Who Bought All Products
# Link: https://leetcode-cn.com/problems/customers-who-bought-all-products

SELECT customer_id
FROM `Customer`
GROUP BY `customer_id`
HAVING COUNT(DISTINCT product_key) = (SELECT COUNT(*) AS cc FROM product)
```

# 1050.actors-and-directors-who-cooperated-at-least-three-times 
```sql
# Title: Actors and Directors Who Cooperated At Least Three Times
# Link: https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times

SELECT `actor_id`, `director_id`
FROM ActorDirector
GROUP BY `actor_id`, `director_id`
HAVING COUNT(*) >= 3
```

# 1068.product-sales-analysis-i 
```sql
# Title: Product Sales Analysis I
# Link: https://leetcode-cn.com/problems/product-sales-analysis-i

SELECT `product_name`, `year`, `price`
FROM `Sales`
         JOIN `Product` ON Sales.product_id = Product.product_id
```

# 1069.product-sales-analysis-ii 
```sql
# Title: Product Sales Analysis II
# Link: https://leetcode-cn.com/problems/product-sales-analysis-ii

SELECT `product_id`, SUM(quantity) AS total_quantity
FROM `sales`
GROUP BY `product_id`
```

# 1070.product-sales-analysis-iii 
```sql


# Title: Product Sales Analysis III
# Link: https://leetcode-cn.com/problems/product-sales-analysis-iii


SELECT `product_id`, `year` AS `first_year`, `quantity`, `price`
FROM `Sales`
WHERE (`product_id`, `year`) IN (SELECT `product_id`, min(`year`)
                                 FROM `Sales`
                                 GROUP BY `product_id`);
```

# 1075.project-employees-i 
```sql
# Title: Project Employees I
# Link: https://leetcode-cn.com/problems/project-employees-i

SELECT `project_id`, round(avg(`experience_years`), 2) AS `average_years`
FROM `Project` AS p
         INNER JOIN `Employee` AS e
                    ON p.employee_id = e.employee_id
GROUP BY `project_id`;
```

# 1076.project-employees-ii 
```sql

# Title: Project Employees II
# Link: https://leetcode-cn.com/problems/project-employees-ii

WITH `tmp` AS (SELECT `project_id`, count(*) AS `c` FROM `project` GROUP BY `project_id`)
SELECT `project_id`
FROM `tmp`
WHERE `c` = (SELECT max(c) FROM `tmp`);
```

# 1082.sales-analysis-i 
```sql
# Title: Sales Analysis I
# Link: https://leetcode-cn.com/problems/sales-analysis-i

SELECT `seller_id`
FROM (SELECT `seller_id`, DENSE_RANK() OVER (ORDER BY total DESC) AS n
      FROM (SELECT `seller_id`, SUM(`price`) AS total FROM `Sales` GROUP BY `seller_id`) t1) t2
WHERE n = 1
```

