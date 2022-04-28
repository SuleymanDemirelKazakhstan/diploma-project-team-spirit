package com.company;


import java.io.*;
import java.sql.*;
import org.postgresql.largeobject.LargeObjectManager;
import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.util.Properties;

public class Main {
    static final String url = "jdbc:postgresql://ec2-3-223-213-207.compute-1.amazonaws.com:5432/d2801cvr39bitk";

    public static void main(String[] args) throws SQLException, IOException {
        setProduct();
        //getProductImage(2);
    }

    public static void setProductWithPath() throws SQLException, IOException {
        Properties props = new Properties();
        props.setProperty("user","wmajjxejdgsbbe");
        props.setProperty("password","d1fa7cf0a11e24417765ed8f97cf93d6688c2b35f58189402e79d29acd44aa60");
//        props.setProperty("ssl","true");
        Connection conn = DriverManager.getConnection(url, props);
        conn.setAutoCommit(false);
        File file = new File("resource/sh1-img10.jpg");
        FileInputStream fis = null;
        fis = new FileInputStream(file);

        PreparedStatement ps = conn.prepareStatement("insert into product" +
                "(shop_id, price, image, name, description, is_auction, discount, crated_at, selled_at)\n" +
                "values(?, ?, ?, ?, ?, ?, ?, ?, ?);");

        ps.setInt(1, 1);
        ps.setInt(2, 6000);
        ps.setString(3, "Блуза");
        ps.setString(4, "Блуза");
        ps.setString(5, "Хлопковая белоснежная блуза! Безумно красивая " +
                "Состояние отличное Состав: 100% хлопок Размер: xs s " +
                "Параметры модели: рост 172, ог 85, от 60, об 90, размер xs-s");
        ps.setBoolean(6, false);
        ps.setDouble(7, 0);
        ps.setDate(8, new Date(new java.util.Date().getTime()));
        ps.setDate(9, null);

        ps.executeUpdate();

        ps.close();
        fis.close();
        conn.commit();
        conn.close();
    }

    public static void setProduct() throws SQLException, IOException {
        Properties props = new Properties();
        props.setProperty("user","wmajjxejdgsbbe");
        props.setProperty("password","d1fa7cf0a11e24417765ed8f97cf93d6688c2b35f58189402e79d29acd44aa60");
//        props.setProperty("ssl","true");
        Connection conn = DriverManager.getConnection(url, props);
        conn.setAutoCommit(false);
        File file = new File("resource/sh1-img10.jpg");
        FileInputStream fis = null;
        fis = new FileInputStream(file);

        PreparedStatement ps = conn.prepareStatement("insert into product" +
                "(shop_id, price, image, name, description, is_auction, discount, crated_at, selled_at)\n" +
                "values(?, ?, ?, ?, ?, ?, ?, ?, ?);");

        ps.setInt(1, 1);
        ps.setInt(2, 6000);
        ps.setBinaryStream(3, fis, (int)file.length());
        ps.setString(4, "Блуза");
        ps.setString(5, "Хлопковая белоснежная блуза! Безумно красивая " +
                "Состояние отличное Состав: 100% хлопок Размер: xs s " +
                "Параметры модели: рост 172, ог 85, от 60, об 90, размер xs-s");
        ps.setBoolean(6, false);
        ps.setDouble(7, 0);
        ps.setDate(8, new Date(new java.util.Date().getTime()));
        ps.setDate(9, null);

        ps.executeUpdate();

        ps.close();
        fis.close();
        conn.commit();
        conn.close();
    }

    public static void getProductImage(int id) throws SQLException, IOException {
        Properties props = new Properties();
        props.setProperty("user","wmajjxejdgsbbe");
        props.setProperty("password","d1fa7cf0a11e24417765ed8f97cf93d6688c2b35f58189402e79d29acd44aa60");
//        props.setProperty("ssl","true");
        Connection conn = DriverManager.getConnection(url, props);
        conn.setAutoCommit(false);


        LargeObjectManager lobj = ((org.postgresql.PGConnection)conn).getLargeObjectAPI();

        PreparedStatement ps = conn.prepareStatement("select * from product where product_id = ?");
        ps.setInt(1, id);
        ResultSet rs = ps.executeQuery();
        while (rs.next()) {
            // Open the large object for reading
            InputStream in = rs.getBinaryStream(4);
            BufferedImage bufferedImage = ImageIO.read(in);
            File outputfile = new File("resource/resultPic" + id + ".jpg");
            ImageIO.write(bufferedImage, "jpg", outputfile);
        }
        rs.close();
        ps.close();
        conn.close();
    }
}
