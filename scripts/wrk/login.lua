wrk.method="POST"
wrk.headers["Content-Type"] = "application/json"
-- 这个要改为你的注册的数据
wrk.body='{"email":"12345@qq.com", "password": "12345@qq.com"}'

-- wrk -t1 -d1s -c2 -s ./scripts/wrk/login.lua http://localhost:8080/users/login