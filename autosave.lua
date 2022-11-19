local bufnr = 27 


vim.api.nvim_create_autocmd("BufWritePost", {
        group = vim.api.nvim_create_augroup("HassanAutoRunner", {clear = true}),
        pattern = "*.go",
        callback = function()
            vim.api.nvim_buf_set_lines(bufnr, 0 , -1, false,{""}) 
            vim.api.nvim_buf_set_lines(bufnr, 0, 0, false, {"output of: main.go"})
            vim.fn.jobstart({"go", "run", "main.go"}, {
                    stdout_buffered = true,
                    on_stdout=function(_, data)
                        if data then
                            vim.api.nvim_buf_set_lines(bufnr, -1 , -1, false, data)
                        end
                    end,
                    on_stderr=function(_, data)
                        if data then
                            vim.api.nvim_buf_set_lines(bufnr, -1 , -1, false, data)
                        end
                    end,
                })
        end,
    })
