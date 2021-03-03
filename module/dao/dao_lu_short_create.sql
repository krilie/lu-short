insert into tb_redirect(id, created_at, updated_at, deleted_at,
                        customer_id, ori_url, `key`, rate_limit, times_limit_left, jump_limit_left, begin_time,
                        end_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
