package com.atguigu.springboot;

import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.web.servlet.support.SpringBootServletInitializer;
import org.springframework.web.servlet.view.UrlBasedViewResolver;

public class ServletInitializer extends SpringBootServletInitializer {

    @Override
    protected SpringApplicationBuilder configure(SpringApplicationBuilder application) {
        // 传入springboot应用的主程序
        return application.sources(SpringBoot04WebJspApplication.class);
    }
    UrlBasedViewResolver

}
