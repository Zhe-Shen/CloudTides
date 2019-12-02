# Generated by Django 2.2.7 on 2019-11-28 01:55

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Policy',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=150, unique=True)),
                ('date_created', models.DateTimeField(blank=True, null=True)),
                ('is_destroy', models.BooleanField(blank=True, null=True)),
                ('deploy_type', models.CharField(choices=[('1', 'kubernetes'), ('2', 'vm')], default='vm', max_length=20)),
                ('idle_policy', models.TextField(blank=True, null=True)),
            ],
            options={
                'verbose_name': 'Tides Policy',
                'verbose_name_plural': 'Tides Policies',
            },
        ),
    ]