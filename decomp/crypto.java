//com.luckyday.app.helpers
public final class RequestHelper {
    public static <T> T decryptString(String str, Type type) {
        Gson create = new GsonBuilder().registerTypeAdapter(UpdateHomePageResponse.class, Deserializers.HOME_PAGE_DESERIALIZER).create();
        try {
            str = DESHelper.decode(str);
            StringBuilder stringBuilder = new StringBuilder();
            stringBuilder.append(type.getClass().getSimpleName());
            stringBuilder.append("String_ ");
            stringBuilder.append(str);
            Log.d("DECODED_STRING", stringBuilder.toString());
            return create.fromJson(str, type);
        } catch (String str2) {
            ThrowableExtension.printStackTrace(str2);
            return null;
        }
    }

    public static <T> EncryptedGsonObject createRequest(T t) {
        try {
            t = new Gson().toJson(t);
            Log.d("ENCODED_STRING", t);
            return new EncryptedGsonObject(DESHelper.encode(t));
        } catch (T t2) {
            ThrowableExtension.printStackTrace(t2);
            return new EncryptedGsonObject("");
        }
    }

    public static String grabCookie(Headers headers) {
        StringBuilder stringBuilder = new StringBuilder();
        for (String str : headers.values("Set-Cookie")) {
            stringBuilder.append(str.substring(0, str.indexOf(";") + 1));
            stringBuilder.append(" ");
        }
        return stringBuilder.toString().trim();
    }
}


//com.luckyday.app.helpers
public final class DESHelper {
    private static final String D_KEY = "DvNw3mJT";
    private static final String E_KEY = "E5QRecZA";

    public static String encode(String str) throws Exception {
        Key generateSecret = SecretKeyFactory.getInstance("DES").generateSecret(new DESKeySpec(E_KEY.getBytes(UrlUtils.UTF8)));
        AlgorithmParameterSpec ivParameterSpec = new IvParameterSpec(E_KEY.getBytes(UrlUtils.UTF8));
        str = str.getBytes(UrlUtils.UTF8);
        Cipher instance = Cipher.getInstance("DES/CBC/PKCS5Padding");
        instance.init(1, generateSecret, ivParameterSpec);
        return Base64.encodeToString(instance.doFinal(str), 0);
    }

    public static String decode(String str) throws InvalidKeyException, NoSuchAlgorithmException, InvalidKeySpecException, NoSuchPaddingException, BadPaddingException, IllegalBlockSizeException, UnsupportedEncodingException, InvalidAlgorithmParameterException {
        Key generateSecret = SecretKeyFactory.getInstance("DES").generateSecret(new DESKeySpec(D_KEY.getBytes(UrlUtils.UTF8)));
        AlgorithmParameterSpec ivParameterSpec = new IvParameterSpec(D_KEY.getBytes(UrlUtils.UTF8));
        Cipher instance = Cipher.getInstance("DES/CBC/PKCS5Padding");
        instance.init(2, generateSecret, ivParameterSpec);
        return new String(instance.doFinal(Base64.decode(str, 0)));
    }
}
