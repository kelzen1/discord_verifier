local verifier_secret_key = "secret_key"
local verifier_server_url = "http://localhost:4000/verify"
local username_discord = common.get_username()
local role = "user_role"
local request_count = 0
local MAX_REQUEST_COUNT = 3
local REQUEST_INTERVAL = 3600 
local last_request_time = 0

function update_loads()
	local time_of = db["time_overflow"] or 0
	time_of = time_of + 1
	db["time_overflow"] = time_of
	db["last_used"] = common.get_unixtime()
end

local verifier = (function(server_url, secret_key)
    local md5 = require "neverlose/md5"

    function generate_signature(username, role, secretKey)
        return md5.sumhexa(string.format("role%susername%s%s",
            role,
            username,
            secretKey
        ))
    end

    function get_code(username, role, callback)
        if request_count >= MAX_REQUEST_COUNT and common.get_unixtime() - last_request_time < REQUEST_INTERVAL then
            print_raw("Too many requests, please try again later.\n\You can request a new code after: " .. REQUEST_INTERVAL .. " seconds")
            return
        end

        if last_request_time ~= 0 and common.get_unixtime() - last_request_time < REQUEST_INTERVAL then
            print_raw("Too many requests, please try again later.\n\You can request a new code after: " .. REQUEST_INTERVAL .. " seconds")
            return
        end
        
        local signature = generate_signature(username, role, secret_key)

        network.post(server_url, {
            ["signature"]   = signature,
            ["username"]    = username,
            ["role"]        = role
        }, {
            ["Content-Type"] = "application/json"
        }, function(result)
            local res = json.parse(result)

            callback(res["success"], res["error"], res["code"])
        end)

        request_count = request_count + 1
        last_request_time = common.get_unixtime()

        update_loads()
    end

    return {
        ["get_code"] = get_code
    }

end)(verifier_server_url, verifier_secret_key)

verifier.get_code(username_discord, role, function(success, err, code)
    print (
        string.format(
            "\nSuccess: %s\nError: %s\nCode: %s",
            success,
            err,
            code
        )
    )
end)
