local verifier_secret_key = "random_secret_password"
local verifier_server_url = "http://localhost:40000/verify"
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

    end

    return {
        ["get_code"] = get_code
    }

end)(verifier_server_url, verifier_secret_key)

verifier.get_code("testuser", "release", function(success, err, code)
    print (
        string.format(
            "\nSuccess: %s\nError: %s\nCode: %s",
            success,
            err,
            code
        )
    )
end)
