package com.atguigu.springboot.controller;

import org.springframework.context.annotation.Configuration;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

/**
 * @ClassName: HelloController
 * @Description: TODO
 * @Author sunsl
 * @Date 2022/2/1 17:20
 * @Version 1.0
 */
@Controller
public class HelloController {

    @GetMapping("/abc")
    public String Hello(Model model) {
        model.addAttribute("msg", "欢迎来到");
        // 返回success.jsp页面
        return "success";
    }
}
