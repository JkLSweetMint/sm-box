insert into
    i18n.languages(code, name, active)
values
    ('ru-RU', 'Русский', true),
    ('en-US', 'English', true),
    ('zh-CN', '中文', true);

select
    i18n.write_text('en-US', 'toasts.error.title', 'An error occured');

select
    i18n.write_text('ru-RU', 'toasts.error.title', 'Произошла ошибка');

select
    i18n.write_text('zh-CN', 'toasts.error.title', '发生错误');

select
    i18n.write_text('en-US', 'auth.form.title', 'Welcome to SM-Box'),
    i18n.write_text('en-US', 'auth.form.description', 'Please, provide your authorization credentials to proceed. '),
    i18n.write_text('en-US', 'auth.form.inputs.username.label', 'Username'),
    i18n.write_text('en-US', 'auth.form.inputs.password.label', 'Password'),
    i18n.write_text('en-US', 'auth.form.buttons.log_in.text', 'Log in'),
    i18n.write_text('en-US', 'auth.form.errors.field_is_required', 'Field is required. '),
    i18n.write_text('en-US', 'auth.form.errors.invalid_value', 'Invalid value. ');

select
    i18n.write_text('ru-RU', 'auth.form.title', 'Добро пожаловать в SM-Box'),
    i18n.write_text('ru-RU', 'auth.form.description', 'Пожалуйста, укажите свои учетные данные для авторизации, чтобы продолжить. '),
    i18n.write_text('ru-RU', 'auth.form.inputs.username.label', 'Имя пользователя'),
    i18n.write_text('ru-RU', 'auth.form.inputs.password.label', 'Пароль'),
    i18n.write_text('ru-RU', 'auth.form.buttons.log_in.text', 'Войти'),
    i18n.write_text('ru-RU', 'auth.form.errors.field_is_required', 'Это поле обязательное. '),
    i18n.write_text('ru-RU', 'auth.form.errors.invalid_value', 'Недопустимое значение. ');

select
    i18n.write_text('zh-CN', 'auth.form.title', '欢迎来到SM-Box'),
    i18n.write_text('zh-CN', 'auth.form.description', '请提供您的登录凭据继续。'),
    i18n.write_text('zh-CN', 'auth.form.inputs.username.label', '用户名称'),
    i18n.write_text('zh-CN', 'auth.form.inputs.password.label', '密码'),
    i18n.write_text('zh-CN', 'auth.form.buttons.log_in.text', '进入'),
    i18n.write_text('zh-CN', 'auth.form.errors.field_is_required', '这个字段是必需的。'),
    i18n.write_text('zh-CN', 'auth.form.errors.invalid_value', '无效值。');

select
    i18n.write_text('en-US', 'auth.project-select.form.title', 'Select a project'),
    i18n.write_text('en-US', 'auth.project-select.form.inputs.select.label', 'Search options...'),
    i18n.write_text('en-US', 'auth.project-select.form.buttons.confirm.text', 'Confirm');

select
    i18n.write_text('ru-RU', 'auth.project-select.form.title', 'Выберите проект'),
    i18n.write_text('ru-RU', 'auth.project-select.form.inputs.select.label', 'Параметры поиска...'),
    i18n.write_text('ru-RU', 'auth.project-select.form.buttons.confirm.text', 'Подтвердить');

select
    i18n.write_text('zh-CN', 'auth.project-select.form.title', '选择项目'),
    i18n.write_text('zh-CN', 'auth.project-select.form.inputs.select.label', '搜索选项。'),
    i18n.write_text('zh-CN', 'auth.project-select.form.buttons.confirm.text', '确认');